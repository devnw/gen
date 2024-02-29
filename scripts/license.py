#!/usr/bin/env python3

import os
import sys
import subprocess

# Define lists of supported file types and directories to ignore
supported = ['go', 'py', 'js', 'ts', 'sh', 'c', 'h', 'cpp', 'hpp', 'rs', 'zig', 'hcl', 'zir']
ignored_dirs = ['scripts', 'vendor', 'node_modules', '.git', '.github', '.vscode', '.idea', '.terraform', '.act', '.circleci', '.gitlab', '.venv', '.cache']

# Function to check if a file should be ignored based on its path
def is_file_in_ignored_dir(file_path, ignored_dirs):
    return any(ignored_dir in file_path for ignored_dir in ignored_dirs)

# Iterate over the command-line arguments (excluding the script name itself)
for file_path in sys.argv[1:]:
    # Extract the file extension
    extension = os.path.splitext(file_path)[1].lstrip('.')

    # Check if the extension is supported
    if extension not in supported:
        continue

    # Check if the file is in an ignored directory
    if is_file_in_ignored_dir(file_path, ignored_dirs):
        continue

    print(f"Adding license to {file_path}")
    # Call the addlicense tool (ensure it's accessible from your Python environment)
    subprocess.run(['addlicense', '-f', './LICENSE_HEADER', '-v', file_path], check=True)
