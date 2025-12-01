#!/usr/bin/env python3
import os
import subprocess

os.chdir('/workspace')

# Read the Makefile and check for TAB vs space characters
with open('Makefile', 'rb') as f:
    content = f.read()

lines = content.split(b'\n')

print("Examining Makefile indentation around lines 70-75:")
for i, line in enumerate(lines[69:75], 70):
    if line.startswith(b' ') or line.startswith(b'\t'):
        # Show the first few characters as hex to see if they're tabs or spaces
        first_chars = line[:10]
        hex_repr = ' '.join(f'{b:02x}' for b in first_chars)
        print(f"Line {i}: {hex_repr} | {line.decode('utf-8', errors='ignore')[:50]}")

print("\nLegend:")
print("09 = TAB character")
print("20 = SPACE character")

# Check if any recipe lines start with TAB
tab_count = 0
space_count = 0
for line in lines:
    if line.startswith(b'\t'):
        tab_count += 1
    elif line.startswith(b'        '):  # 8 spaces
        space_count += 1

print(f"\nSummary:")
print(f"Lines starting with TAB: {tab_count}")
print(f"Lines starting with 8 spaces: {space_count}")

# Test make syntax
print("\nTesting make syntax:")
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