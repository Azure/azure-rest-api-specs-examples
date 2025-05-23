from azure.identity import DefaultAzureCredential

from azure.mgmt.network import NetworkManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-network
# USAGE
    python route_map_put.py

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

    response = client.route_maps.begin_create_or_update(
        resource_group_name="rg1",
        virtual_hub_name="virtualHub1",
        route_map_name="routeMap1",
        route_map_parameters={
            "properties": {
                "associatedInboundConnections": [
                    "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/expressRouteGateways/exrGateway1/expressRouteConnections/exrConn1"
                ],
                "associatedOutboundConnections": [],
                "rules": [
                    {
                        "actions": [
                            {"parameters": [{"asPath": ["22334"], "community": [], "routePrefix": []}], "type": "Add"}
                        ],
                        "matchCriteria": [
                            {"asPath": [], "community": [], "matchCondition": "Contains", "routePrefix": ["10.0.0.0/8"]}
                        ],
                        "name": "rule1",
                        "nextStepIfMatched": "Continue",
                    }
                ],
            }
        },
    ).result()
    print(response)


# x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/RouteMapPut.json
if __name__ == "__main__":
    main()
