#!/bin/bash
cd /workspace

echo "Testing Makefile syntax..."
# Try to run make with dry-run to test syntax
if command -v make >/dev/null 2>&1; then
    echo "Make is available, testing syntax:"
    make -n test 2>&1 | head -10
else
    echo "Make command not found"
fi

echo ""
echo "Checking raw bytes of line 70:"
sed -n '70p' Makefile | hexdump -C | head -1

echo ""
echo "Checking if line 70 starts with TAB (0x09):"
sed -n '70p' Makefile | od -c | head -1