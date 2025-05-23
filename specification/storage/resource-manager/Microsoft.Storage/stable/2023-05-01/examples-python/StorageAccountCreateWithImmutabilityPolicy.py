from azure.identity import DefaultAzureCredential

from azure.mgmt.storage import StorageManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-storage
# USAGE
    python storage_account_create_with_immutability_policy.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = StorageManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="{subscription-id}",
    )

    response = client.storage_accounts.begin_create(
        resource_group_name="res9101",
        account_name="sto4445",
        parameters={
            "extendedLocation": {"name": "losangeles001", "type": "EdgeZone"},
            "kind": "Storage",
            "location": "eastus",
            "properties": {
                "immutableStorageWithVersioning": {
                    "enabled": True,
                    "immutabilityPolicy": {
                        "allowProtectedAppendWrites": True,
                        "immutabilityPeriodSinceCreationInDays": 15,
                        "state": "Unlocked",
                    },
                }
            },
            "sku": {"name": "Standard_GRS"},
        },
    ).result()
    print(response)


# x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2023-05-01/examples/StorageAccountCreateWithImmutabilityPolicy.json
if __name__ == "__main__":
    main()
