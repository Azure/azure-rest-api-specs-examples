from azure.identity import DefaultAzureCredential
from azure.mgmt.connectedvmware import ConnectedVMwareMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-connectedvmware
# USAGE
    python delete_resource_pool.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ConnectedVMwareMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="fd3c3665-1729-4b7b-9a38-238e83b0f98b",
    )

    client.resource_pools.begin_delete(
        resource_group_name="testrg",
        resource_pool_name="HRPool",
    ).result()


# x-ms-original-file: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/stable/2023-10-01/examples/DeleteResourcePool.json
if __name__ == "__main__":
    main()
