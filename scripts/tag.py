#!/usr/bin/env python3

import subprocess
import sys

def run_command(command, capture_output=True, text=True):
    """Executes a shell command and optionally returns its output."""
    result = subprocess.run(command, capture_output=capture_output, text=text, shell=True)
    if result.returncode != 0:
        print(f"Command '{command}' failed with return code {result.returncode}")
        sys.exit(result.returncode)
    out = result.stdout
    try:
        return out.strip()
    except AttributeError:
        return out

def main():
    # Fetch the latest tag
    latest_tag = run_command("git describe --tags $(git rev-list --tags --max-count=1)")
    current_major, current_minor, current_patch = latest_tag.split('.')

    # Calculate the next minor version
    next_minor = str(int(current_minor) + 1)
    default_version = f"{current_major}.{next_minor}.{current_patch}"

    # Prompt the user for a version number, defaulting to the next minor version
    version = input(f"Enter the version number [{default_version}]: ").strip() or default_version

    # Collect commits since the last tag
    commits = run_command(f"git log {latest_tag}..HEAD --pretty=format:'%h %s'", capture_output=True, text=True)
    commits_formatted = "\n".join([f"- {line}" for line in commits.split('\n')])

    # Create and push the new tag
    tag_message = f"Release {version}\n\n{commits_formatted}"
    run_command(f"git tag -a {version} -m \"{tag_message}\"", capture_output=False, text=False)
    run_command(f"git push origin {version}", capture_output=False, text=False)

if __name__ == "__main__":
    main()
