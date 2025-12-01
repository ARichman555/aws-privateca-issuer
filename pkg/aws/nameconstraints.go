/*
Copyright 2021 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package aws

import (
        "crypto/x509"
        "crypto/x509/pkix"
        "encoding/asn1"
        "encoding/pem"
        "fmt"

        acmpcatypes "github.com/aws/aws-sdk-go-v2/service/acmpca/types"
        api "github.com/cert-manager/aws-privateca-issuer/pkg/api/v1beta1"
        cmapi "github.com/cert-manager/cert-manager/pkg/apis/certmanager/v1"
)

// Name Constraints OID as defined in RFC 5280
var nameConstraintsOID = asn1.ObjectIdentifier{2, 5, 29, 30}

// NameConstraints represents the ASN.1 structure for name constraints extension
type NameConstraints struct {
        PermittedSubtrees []GeneralSubtree `asn1:"optional,tag:0"`
        ExcludedSubtrees  []GeneralSubtree `asn1:"optional,tag:1"`
}

// GeneralSubtree represents a general subtree in name constraints
type GeneralSubtree struct {
        Base    GeneralName `asn1:""`
        Minimum int         `asn1:"optional,tag:0,default:0"`
        Maximum int         `asn1:"optional,tag:1"`
}

// GeneralName represents a general name in name constraints
type GeneralName struct {
        DNSName string `asn1:"tag:2,optional"`
        // Add other name types as needed
}

// hasNameConstraints checks if a CertificateRequest contains name constraints extension
func hasNameConstraints(cr *cmapi.CertificateRequest) (bool, error) {
        block, _ := pem.Decode(cr.Spec.Request)
        if block == nil {
                return false, fmt.Errorf("failed to decode CSR")
        }

        csr, err := x509.ParseCertificateRequest(block.Bytes)
        if err != nil {
                return false, fmt.Errorf("failed to parse CSR: %v", err)
        }

        // Check if name constraints extension is present
        for _, ext := range csr.Extensions {
                if ext.Id.Equal(nameConstraintsOID) {
                        return true, nil
                }
        }

        return false, nil
}

// extractNameConstraints extracts name constraints from a CertificateRequest
func extractNameConstraints(cr *cmapi.CertificateRequest) (*NameConstraints, error) {
        block, _ := pem.Decode(cr.Spec.Request)
        if block == nil {
                return nil, fmt.Errorf("failed to decode CSR")
        }

        csr, err := x509.ParseCertificateRequest(block.Bytes)
        if err != nil {
                return nil, fmt.Errorf("failed to parse CSR: %v", err)
        }

        // Find name constraints extension
        for _, ext := range csr.Extensions {
                if ext.Id.Equal(nameConstraintsOID) {
                        var nc NameConstraints
                        if _, err := asn1.Unmarshal(ext.Value, &nc); err != nil {
                                return nil, fmt.Errorf("failed to parse name constraints: %v", err)
                        }
                        return &nc, nil
                }
        }

        return nil, fmt.Errorf("name constraints extension not found")
}

// buildApiPassthrough constructs the ApiPassthrough structure for name constraints
func buildApiPassthrough(nc *NameConstraints) (*acmpcatypes.ApiPassthrough, error) {
        // Encode the name constraints extension
        ncBytes, err := asn1.Marshal(*nc)
        if err != nil {
                return nil, fmt.Errorf("failed to marshal name constraints: %v", err)
        }

        // Create the extension
        ext := acmpcatypes.Extensions{
                ExtnId:    nameConstraintsOID.String(),
                Critical:  true,
                ExtnValue: ncBytes,
        }

        // Build the ApiPassthrough structure
        apiPassthrough := &acmpcatypes.ApiPassthrough{
                Extensions: &acmpcatypes.Extensions{
                        CertificatePolicies:   ext.CertificatePolicies,
                        ExtendedKeyUsage:      ext.ExtendedKeyUsage,
                        KeyUsage:              ext.KeyUsage,
                        SubjectAlternativeName: ext.SubjectAlternativeName,
                        CustomExtensions: []acmpcatypes.CustomExtension{
                                {
                                        ObjectIdentifier: nameConstraintsOID.String(),
                                        Value:           ncBytes,
                                        Critical:        true,
                                },
                        },
                },
        }

        return apiPassthrough, nil
}