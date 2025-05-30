from azure.identity import DefaultAzureCredential

from azure.mgmt.network import NetworkManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-network
# USAGE
    python express_route_circuit_create_on_express_route_port.py

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

    response = client.express_route_circuits.begin_create_or_update(
        resource_group_name="rg1",
        circuit_name="expressRouteCircuit1",
        parameters={
            "location": "westus",
            "properties": {
                "authorizationKey": "b0be57f5-1fba-463b-adec-ffe767354cdd",
                "bandwidthInGbps": 10,
                "enableDirectPortRateLimit": False,
                "expressRoutePort": {
                    "id": "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/expressRoutePorts/portName"
                },
            },
            "sku": {"family": "MeteredData", "name": "Premium_MeteredData", "tier": "Premium"},
        },
    ).result()
    print(response)


# x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-05-01/examples/ExpressRouteCircuitCreateOnExpressRoutePort.json
if __name__ == "__main__":
    main()
