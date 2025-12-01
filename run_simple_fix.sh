#!/bin/bash
cd /workspace
python3 simple_fix.py
echo "Fix completed. Checking result:"
echo "Line 70 after fix:"
sed -n '70p' Makefile | cat -A