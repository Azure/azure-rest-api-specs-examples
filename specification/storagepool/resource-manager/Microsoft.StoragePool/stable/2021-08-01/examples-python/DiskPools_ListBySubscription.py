from azure.identity import DefaultAzureCredential
from azure.mgmt.storagepool import StoragePoolManagement

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-storagepool
# USAGE
    python list_disk_pools_by_subscription.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = StoragePoolManagement(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-0000-0000-0000-000000000000",
    )

    response = client.disk_pools.list_by_subscription()
    for item in response:
        print(item)


# x-ms-original-file: specification/storagepool/resource-manager/Microsoft.StoragePool/stable/2021-08-01/examples/DiskPools_ListBySubscription.json
if __name__ == "__main__":
    main()
