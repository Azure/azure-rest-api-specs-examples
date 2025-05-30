from azure.identity import DefaultAzureCredential

from azure.mgmt.storage import StorageManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-storage
# USAGE
    python queue_operation_get.py

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

    response = client.queue.get(
        resource_group_name="res3376",
        account_name="sto328",
        queue_name="queue6185",
    )
    print(response)


# x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2024-01-01/examples/QueueOperationGet.json
if __name__ == "__main__":
    main()
