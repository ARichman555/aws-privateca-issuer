#!/bin/bash
# Check if Makefile has TAB characters
cd /workspace
echo "Checking for TAB characters in Makefile around line 70:"
sed -n '70,72p' Makefile | cat -A
echo ""
echo "Checking if make command works:"
make --version > /dev/null 2>&1 && echo "Make is available" || echo "Make not available"
echo ""
echo "Testing Makefile syntax:"
make -n test 2>&1 | head -5