from azure.identity import DefaultAzureCredential

from azure.mgmt.network import NetworkManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-network
# USAGE
    python express_route_circuit_connection_list.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = NetworkManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="subid1",
    )

    response = client.express_route_circuit_connections.list(
        resource_group_name="rg1",
        circuit_name="ExpressRouteARMCircuitA",
        peering_name="AzurePrivatePeering",
    )
    for item in response:
        print(item)


# x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/ExpressRouteCircuitConnectionList.json
if __name__ == "__main__":
    main()
