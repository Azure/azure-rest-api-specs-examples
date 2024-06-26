from typing import Any, IO, Union

from azure.identity import DefaultAzureCredential

from azure.mgmt.storagecache import StorageCacheManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-storagecache
# USAGE
    python aml_filesystems_update.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = StorageCacheManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-0000-0000-0000-000000000000",
    )

    response = client.aml_filesystems.begin_update(
        resource_group_name="scgroup",
        aml_filesystem_name="fs1",
        aml_filesystem={
            "properties": {
                "encryptionSettings": {
                    "keyEncryptionKey": {
                        "keyUrl": "https://examplekv.vault.azure.net/keys/kvk/3540a47df75541378d3518c6a4bdf5af",
                        "sourceVault": {
                            "id": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scgroup/providers/Microsoft.KeyVault/vaults/keyvault-cmk"
                        },
                    }
                },
                "maintenanceWindow": {"dayOfWeek": "Friday", "timeOfDayUTC": "22:00"},
                "rootSquashSettings": {
                    "mode": "All",
                    "noSquashNidLists": "10.0.0.[5-6]@tcp;10.0.1.2@tcp",
                    "squashGID": 99,
                    "squashUID": 99,
                },
            },
            "tags": {"Dept": "ContosoAds"},
        },
    ).result()
    print(response)


# x-ms-original-file: specification/storagecache/resource-manager/Microsoft.StorageCache/stable/2024-03-01/examples/amlFilesystems_Update.json
if __name__ == "__main__":
    main()
