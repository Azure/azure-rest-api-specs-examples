from azure.identity import DefaultAzureCredential

from azure.mgmt.storagediscovery import StorageDiscoveryMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-storagediscovery
# USAGE
    python storage_discovery_workspaces_create_or_update.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = StorageDiscoveryMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.storage_discovery_workspaces.create_or_update(
        resource_group_name="sample-rg",
        storage_discovery_workspace_name="Sample-Storage-Workspace",
        resource={
            "location": "westeurope",
            "properties": {
                "description": "Sample Storage Discovery Workspace",
                "scopes": [
                    {
                        "displayName": "Sample-Collection",
                        "resourceTypes": [
                            "/subscriptions/b79cb3ba-745e-5d9a-8903-4a02327a7e09/resourceGroups/sample-rg/providers/Microsoft.Storage/storageAccounts/sample-storageAccount"
                        ],
                        "tagKeysOnly": ["filterTag1", "filterTag2"],
                        "tags": {"filterTag3": "value3", "filterTag4": "value4"},
                    },
                    {
                        "displayName": "Sample-Collection-2",
                        "resourceTypes": [
                            "/subscriptions/b79cb3ba-745e-5d9a-8903-4a02327a7e09/resourceGroups/sample-rg/providers/Microsoft.Storage/storageAccounts/sample-storageAccount"
                        ],
                        "tagKeysOnly": ["filterTag5"],
                        "tags": {"filterTag6": "value6"},
                    },
                ],
                "workspaceRoots": ["/subscriptions/b79cb3ba-745e-5d9a-8903-4a02327a7e09"],
            },
            "tags": {"tag1": "value1", "tag2": "value2"},
        },
    )
    print(response)


# x-ms-original-file: 2025-06-01-preview/StorageDiscoveryWorkspaces_CreateOrUpdate.json
if __name__ == "__main__":
    main()
