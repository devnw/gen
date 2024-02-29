#!/usr/bin/env python3

import os
import subprocess
import sys

def run_command(command):
    """Executes a shell command and returns its output."""
    return subprocess.check_output(command, shell=True, text=True)

def main():
    fuzz_time = os.environ.get('FUZZ_TIME', '10')

    try:
        # Find all Go test files containing fuzz functions
        files = run_command("grep -r --include='**_test.go' --files-with-matches 'func Fuzz' .").strip().split('\n')

        for file in files:
            # Extract fuzz function names from each file
            funcs = run_command(f"grep -o 'func Fuzz\\w*' {file} | sed 's/func //'").strip().split('\n')

            for func in funcs:
                print(f"Fuzzing {func} in {file}")
                parent_dir = os.path.dirname(file)

                # Run the fuzz function with go test
                try:
                    subprocess.check_call(f"go test {parent_dir} -run={func} -fuzz={func} -fuzztime={fuzz_time}s", shell=True)
                except subprocess.CalledProcessError:
                    print(f"Fuzzing {func} in {file} failed")
                    sys.exit(1)
    except subprocess.CalledProcessError as e:
        # Handle cases where the initial grep command finds no matching files
        print("No Go test files with fuzz functions found.")
    except Exception as e:
        print(f"An error occurred: {e}")
        sys.exit(1)

if __name__ == "__main__":
    main()
