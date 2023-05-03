import json
import sys

if len(sys.argv) != 3:
    print("Usage: python program.py <file1.json> <file2.json>")
    sys.exit(1)

# Load the contents of file1 into memory
with open(sys.argv[1], "r") as file1:
    data1 = json.load(file1)

# Load the contents of file2 into memory
with open(sys.argv[2], "r") as file2:
    data2 = json.load(file2)

# Find the service account block to replace based on the account_id field
account_id = "my-service-account-4"  # Replace this with the account ID you want to replace
account_index = None
for i, resource in enumerate(data1["resource"]):
    if resource["type"] == "google_service_account" and resource["properties"]["account_id"] == account_id:
        account_index = i
        break

# If a matching service account block is found, replace it with the corresponding block from file2
if account_index is not None:
    account_found = False
    for resource in data2["resource"]:
        if resource["type"] == "google_service_account" and resource["properties"]["account_id"] == account_id:
            data1["resource"][account_index] = resource
            account_found = True
            break
    if not account_found:
        print(f"Error: Account with ID '{account_id}' not found in '{sys.argv[2]}'")
else:
    print(f"Error: Account with ID '{account_id}' not found in '{sys.argv[1]}'")

# Write the updated contents to a new file called result.json
with open("result.json", "w") as result_file:
    json.dump(data1, result_file, indent=2)

