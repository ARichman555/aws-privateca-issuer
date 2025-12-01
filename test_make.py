#!/usr/bin/env python3
import subprocess
import os

os.chdir('/workspace')

# Test if make command works with current Makefile
try:
    result = subprocess.run(['make', '--dry-run', 'test'], 
                          capture_output=True, text=True, timeout=10)
    print("Make command exit code:", result.returncode)
    print("STDOUT:", result.stdout[:500])
    print("STDERR:", result.stderr[:500])
except Exception as e:
    print("Error running make:", e)

# Check for TAB characters in the Makefile
print("\nChecking for TAB characters around line 70:")
with open('Makefile', 'rb') as f:
    lines = f.readlines()
    for i, line in enumerate(lines[69:72], 70):
        print(f"Line {i}: {repr(line)}")