from azure.identity import DefaultAzureCredential
from azure.mgmt.managednetworkfabric import ManagedNetworkFabricMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-managednetworkfabric
# USAGE
    python network_to_network_interconnects_update_maximum_set_gen.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ManagedNetworkFabricMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="1234ABCD-0A1B-1234-5678-123456ABCDEF",
    )

    response = client.network_to_network_interconnects.begin_update(
        resource_group_name="example-rg",
        network_fabric_name="example-fabric",
        network_to_network_interconnect_name="example-nni",
        body={
            "properties": {
                "egressAclId": "/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/accessControlLists/example-acl",
                "exportRoutePolicy": {
                    "exportIpv4RoutePolicyId": "/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/routePolicies/example-routePolicy1",
                    "exportIpv6RoutePolicyId": "/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/routePolicies/example-routePolicy1",
                },
                "importRoutePolicy": {
                    "importIpv4RoutePolicyId": "/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/routePolicies/example-routePolicy1",
                    "importIpv6RoutePolicyId": "/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/routePolicies/example-routePolicy1",
                },
                "ingressAclId": "/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/accessControlLists/example-acl",
                "layer2Configuration": {
                    "interfaces": [
                        "/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkDevices/example-networkDevice/networkInterfaces/example-networkInterface"
                    ],
                    "mtu": 1500,
                },
                "npbStaticRouteConfiguration": {
                    "bfdConfiguration": {"intervalInMilliSeconds": 310, "multiplier": 15},
                    "ipv4Routes": [{"nextHop": ["21.20.20.10"], "prefix": "20.0.0.11/30"}],
                    "ipv6Routes": [{"nextHop": ["5FFE:FFFF:0:CD30::ac"], "prefix": "4FFE:FFFF:0:CD30::ac/127"}],
                },
                "optionBLayer3Configuration": {
                    "peerASN": 2345,
                    "primaryIpv4Prefix": "20.0.0.12/29",
                    "primaryIpv6Prefix": "4FFE:FFFF:0:CD30::a8/127",
                    "secondaryIpv4Prefix": "20.0.0.14/29",
                    "secondaryIpv6Prefix": "6FFE:FFFF:0:CD30::ac/127",
                    "vlanId": 1235,
                },
            }
        },
    ).result()
    print(response)


# x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/NetworkToNetworkInterconnects_Update_MaximumSet_Gen.json
if __name__ == "__main__":
    main()
