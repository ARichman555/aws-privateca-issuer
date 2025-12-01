#!/usr/bin/env python3
import os

os.chdir('/workspace')

# Read Makefile and check line 70 specifically
with open('Makefile', 'rb') as f:
    lines = f.readlines()

print("Checking line 70 (the one mentioned in the error):")
if len(lines) >= 70:
    line70 = lines[69]  # 0-indexed
    print(f"Line 70 raw bytes: {line70}")
    print(f"Line 70 hex: {line70.hex()}")
    
    # Check if it starts with TAB (0x09) or spaces (0x20)
    if line70.startswith(b'\t'):
        print("✓ Line 70 starts with TAB character")
    elif line70.startswith(b'        '):  # 8 spaces
        print("✗ Line 70 starts with 8 spaces (should be TAB)")
    else:
        print(f"? Line 70 starts with: {repr(line70[:10])}")

# Check a few more lines around it
print("\nChecking lines 69-72:")
for i in range(68, min(72, len(lines))):
    line = lines[i]
    line_num = i + 1
    if line.strip():  # Only check non-empty lines
        if line.startswith(b'\t'):
            status = "✓ TAB"
        elif line.startswith(b'        '):
            status = "✗ 8 SPACES"
        elif line.startswith(b' '):
            status = f"? {len(line) - len(line.lstrip())} spaces"
        else:
            status = "- No indent"
        print(f"Line {line_num}: {status} | {line.decode('utf-8', errors='ignore').rstrip()[:50]}")

# Test make syntax
print("\nTesting make syntax:")
import subprocess
try:
    result = subprocess.run(['make', '--version'], capture_output=True, text=True, timeout=5)
    if result.returncode == 0:
        print("Make is available")
        # Test the actual syntax
        result = subprocess.run(['make', '-n', 'test'], capture_output=True, text=True, timeout=10)
        if result.returncode == 0:
            print("✓ Makefile syntax is correct")
        else:
            print(f"✗ Makefile syntax error: {result.stderr}")
    else:
        print("Make command failed")
except Exception as e:
    print(f"Error testing make: {e}")