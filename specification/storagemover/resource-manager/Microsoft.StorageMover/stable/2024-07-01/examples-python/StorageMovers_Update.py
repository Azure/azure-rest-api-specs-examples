from typing import Any, IO, Union

from azure.identity import DefaultAzureCredential

from azure.mgmt.storagemover import StorageMoverMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-storagemover
# USAGE
    python storage_movers_update.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = StorageMoverMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="60bcfc77-6589-4da2-b7fd-f9ec9322cf95",
    )

    response = client.storage_movers.update(
        resource_group_name="examples-rg",
        storage_mover_name="examples-storageMoverName",
        storage_mover={"properties": {"description": "Updated Storage Mover Description"}},
    )
    print(response)


# x-ms-original-file: specification/storagemover/resource-manager/Microsoft.StorageMover/stable/2024-07-01/examples/StorageMovers_Update.json
if __name__ == "__main__":
    main()
