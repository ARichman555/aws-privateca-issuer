package testutil

// Common test value builders to reduce duplication

// ServiceValues creates common service configuration values
func ServiceValues(serviceType string, port int) map[string]interface{} {
	return map[string]interface{}{
		"service": map[string]interface{}{
			"type": serviceType,
			"port": port,
		},
	}
}

// ImageValues creates image configuration values
func ImageValues(repository, tag, pullPolicy string) map[string]interface{} {
	return map[string]interface{}{
		"image": map[string]interface{}{
			"repository": repository,
			"tag":        tag,
			"pullPolicy": pullPolicy,
		},
	}
}

// ServiceAccountValues creates service account configuration values
func ServiceAccountValues(name string, annotations map[string]string) map[string]interface{} {
	values := map[string]interface{}{
		"serviceAccount": map[string]interface{}{
			"create": true,
		},
	}
	
	if name != "" {
		values["serviceAccount"].(map[string]interface{})["name"] = name
	}
	
	if len(annotations) > 0 {
		values["serviceAccount"].(map[string]interface{})["annotations"] = annotations
	}
	
	return values
}

// RBACValues creates RBAC configuration values
func RBACValues(enabled bool) map[string]interface{} {
	return map[string]interface{}{
		"rbac": map[string]interface{}{
			"create": enabled,
		},
		"serviceAccount": map[string]interface{}{
			"create": enabled,
		},
	}
}

// NameOverrideValues creates name override values
func NameOverrideValues(nameOverride, fullnameOverride string) map[string]interface{} {
	values := make(map[string]interface{})
	
	if nameOverride != "" {
		values["nameOverride"] = nameOverride
	}
	
	if fullnameOverride != "" {
		values["fullnameOverride"] = fullnameOverride
	}
	
	return values
}
