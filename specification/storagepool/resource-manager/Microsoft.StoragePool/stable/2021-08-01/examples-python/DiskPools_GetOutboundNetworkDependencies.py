from azure.identity import DefaultAzureCredential
from azure.mgmt.storagepool import StoragePoolManagement

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-storagepool
# USAGE
    python get_disk_pool_outbound_network_dependencies.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = StoragePoolManagement(
        credential=DefaultAzureCredential(),
        subscription_id="34adfa4f-cedf-4dc0-ba29-b6d1a69ab345",
    )

    response = client.disk_pools.list_outbound_network_dependencies_endpoints(
        resource_group_name="Sample-WestUSResourceGroup",
        disk_pool_name="SampleAse",
    )
    for item in response:
        print(item)


# x-ms-original-file: specification/storagepool/resource-manager/Microsoft.StoragePool/stable/2021-08-01/examples/DiskPools_GetOutboundNetworkDependencies.json
if __name__ == "__main__":
    main()
