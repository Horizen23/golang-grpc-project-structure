#!/bin/bash

# Extract current module name from go.mod
old_module_name=$(grep "^module" go.mod | awk '{print $2}')

# Check if the current module name is found
if [ -z "$old_module_name" ]; then
    echo "Error: Failed to extract current module name from go.mod"
    exit 1
fi

# Check if the correct number of arguments is provided
if [ "$#" -ne 1 ]; then
    echo "Usage: $0 <new_module_name>"
    exit 1
fi

new_module_name="$1"

# Update module name in go.mod
sed -i '' "s/$old_module_name/$new_module_name/g" go.mod

# Replace old module name with new module name in Go files
find . -name '*.go' -type f -exec sed -i '' "s/$old_module_name/$new_module_name/g" {} +

echo "Module name updated from $old_module_name to $new_module_name"
