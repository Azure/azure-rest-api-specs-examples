from azure.identity import DefaultAzureCredential

from azure.mgmt.storage import StorageManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-storage
# USAGE
    python blob_ranges_restore.py

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

    response = client.storage_accounts.begin_restore_blob_ranges(
        resource_group_name="res9101",
        account_name="sto4445",
        parameters={
            "blobRanges": [
                {"endRange": "container/blobpath2", "startRange": "container/blobpath1"},
                {"endRange": "", "startRange": "container2/blobpath3"},
            ],
            "timeToRestore": "2019-04-20T15:30:00.0000000Z",
        },
    ).result()
    print(response)


# x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2023-05-01/examples/BlobRangesRestore.json
if __name__ == "__main__":
    main()
