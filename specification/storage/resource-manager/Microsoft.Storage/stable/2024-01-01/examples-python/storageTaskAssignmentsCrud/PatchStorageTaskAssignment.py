from azure.identity import DefaultAzureCredential

from azure.mgmt.storage import StorageManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-storage
# USAGE
    python patch_storage_task_assignment.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = StorageManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="1f31ba14-ce16-4281-b9b4-3e78da6e1616",
    )

    response = client.storage_task_assignments.begin_update(
        resource_group_name="res4228",
        account_name="sto4445",
        storage_task_assignment_name="myassignment1",
        parameters={
            "properties": {
                "description": "My Storage task assignment",
                "enabled": True,
                "executionContext": {
                    "target": {"excludePrefix": [], "prefix": ["prefix1", "prefix2"]},
                    "trigger": {"parameters": {"startOn": "2022-11-15T21:52:47.8145095Z"}, "type": "RunOnce"},
                },
                "report": {"prefix": "container1"},
            }
        },
    ).result()
    print(response)


# x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2024-01-01/examples/storageTaskAssignmentsCrud/PatchStorageTaskAssignment.json
if __name__ == "__main__":
    main()
