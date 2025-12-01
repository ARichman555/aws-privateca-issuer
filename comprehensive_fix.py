#!/usr/bin/env python3
import os
import subprocess

os.chdir('/workspace')

# Read the current Makefile
with open('Makefile', 'r') as f:
    lines = f.readlines()

print("Fixing Makefile indentation and golangci-lint integration...")

# Create backup
with open('Makefile.backup', 'w') as f:
    f.writelines(lines)

print("Created backup: Makefile.backup")

# Fix each line
fixed_lines = []
changes_made = 0

for i, line in enumerate(lines):
    # Check if line starts with exactly 8 spaces (recipe line)
    if line.startswith('        ') and not line.startswith('         '):  # 8 spaces, not 9+
        # Replace the 8 spaces with a single TAB
        fixed_line = '\t' + line[8:]
        fixed_lines.append(fixed_line)
        changes_made += 1
        print(f"Fixed line {i+1}: spaces -> TAB")
    else:
        fixed_lines.append(line)

# Write the fixed content
with open('Makefile', 'w') as f:
    f.writelines(fixed_lines)

print(f"Fixed {changes_made} recipe lines with proper TAB indentation")

# Now fix the lint target and golangci-lint version
with open('Makefile', 'r') as f:
    content = f.read()

# Fix lint target
old_lint = '''lint:
\techo "Linter is deprecated with go1.18!"'''

new_lint = '''lint: golangci-lint
\t$(GOLANGCILINT) run --timeout 10m'''

if old_lint in content:
    content = content.replace(old_lint, new_lint)
    print("Updated lint target to use golangci-lint")
else:
    print("Lint target pattern not found, checking alternative...")
    # Try alternative pattern
    if 'echo "Linter is deprecated with go1.18!"' in content:
        content = content.replace(
            'lint:\n\techo "Linter is deprecated with go1.18!"',
            'lint: golangci-lint\n\t$(GOLANGCILINT) run --timeout 10m'
        )
        print("Updated lint target (alternative pattern)")

# Update golangci-lint version
old_version = 'github.com/golangci/golangci-lint/cmd/golangci-lint@v1.35.2'
new_version = 'github.com/golangci/golangci-lint/cmd/golangci-lint@v1.55.2'

if old_version in content:
    content = content.replace(old_version, new_version)
    print("Updated golangci-lint version to v1.55.2")
else:
    print("golangci-lint version not found or already updated")

# Write the final content
with open('Makefile', 'w') as f:
    f.write(content)

print("Makefile has been updated with proper TAB indentation and golangci-lint integration")

# Test the syntax
print("\nTesting Makefile syntax...")
try:
    result = subprocess.run(['make', '--dry-run', 'test'], capture_output=True, text=True, timeout=10)
    if result.returncode == 0:
        print("✓ Makefile syntax is now correct!")
        print("✓ CI pipeline should now pass the make test command")
    else:
        print(f"✗ Makefile syntax error: {result.stderr}")
        print("Rolling back changes...")
        # Restore backup
        with open('Makefile.backup', 'r') as f:
            backup_content = f.read()
        with open('Makefile', 'w') as f:
            f.write(backup_content)
        print("Restored from backup")
except Exception as e:
    print(f"Error testing make: {e}")