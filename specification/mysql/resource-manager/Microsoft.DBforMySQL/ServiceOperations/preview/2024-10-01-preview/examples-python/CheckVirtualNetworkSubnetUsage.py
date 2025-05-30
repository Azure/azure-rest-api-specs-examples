from azure.identity import DefaultAzureCredential

from azure.mgmt.mysqlflexibleservers import MySQLManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-mysqlflexibleservers
# USAGE
    python check_virtual_network_subnet_usage.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = MySQLManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="ffffffff-ffff-ffff-ffff-ffffffffffff",
    )

    response = client.check_virtual_network_subnet_usage.execute(
        location_name="WestUS",
        parameters={
            "virtualNetworkResourceId": "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.Network/virtualNetworks/testvnet"
        },
    )
    print(response)


# x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/ServiceOperations/preview/2024-10-01-preview/examples/CheckVirtualNetworkSubnetUsage.json
if __name__ == "__main__":
    main()
