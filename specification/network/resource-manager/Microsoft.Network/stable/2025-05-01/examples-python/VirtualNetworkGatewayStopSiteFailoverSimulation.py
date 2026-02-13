from azure.identity import DefaultAzureCredential

from azure.mgmt.network import NetworkManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-network
# USAGE
    python virtual_network_gateway_stop_site_failover_simulation.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = NetworkManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="subid",
    )

    response = client.virtual_network_gateways.begin_stop_express_route_site_failover_simulation(
        resource_group_name="rg1",
        virtual_network_gateway_name="ergw",
        stop_parameters={
            "details": [
                {"failoverConnectionName": "conn1", "failoverLocation": "Denver", "isVerified": False},
                {"failoverConnectionName": "conn2", "failoverLocation": "Amsterdam", "isVerified": True},
            ],
            "peeringLocation": "Vancouver",
            "wasSimulationSuccessful": True,
        },
    ).result()
    print(response)


# x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/VirtualNetworkGatewayStopSiteFailoverSimulation.json
if __name__ == "__main__":
    main()
