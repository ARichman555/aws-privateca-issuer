#!/bin/bash
# Test if the Makefile has correct TAB indentation
cd /workspace
make --dry-run test 2>&1 | head -20