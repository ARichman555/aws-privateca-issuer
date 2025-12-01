#!/usr/bin/env python3
import subprocess
import os

os.chdir('/workspace')

# Try to restore Makefile from git
try:
    result = subprocess.run(['git', 'checkout', 'HEAD', '--', 'Makefile'], 
                          capture_output=True, text=True, timeout=10)
    if result.returncode == 0:
        print("✓ Restored Makefile from git")
    else:
        print(f"✗ Git restore failed: {result.stderr}")
        print("Trying git status...")
        status_result = subprocess.run(['git', 'status'], capture_output=True, text=True, timeout=5)
        print(f"Git status: {status_result.stdout}")
except Exception as e:
    print(f"Error with git: {e}")

# Check if Makefile exists now
if os.path.exists('Makefile'):
    print("✓ Makefile exists")
    with open('Makefile', 'r') as f:
        lines = f.readlines()
        print(f"Makefile has {len(lines)} lines")
        if len(lines) > 69:
            print(f"Line 70: {repr(lines[69])}")
else:
    print("✗ Makefile still missing")