from azure.identity import DefaultAzureCredential
from azure.mgmt.managednetworkfabric import ManagedNetworkFabricMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-managednetworkfabric
# USAGE
    python network_fabrics_create_maximum_set_gen.py

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

    response = client.network_fabrics.begin_create(
        resource_group_name="example-rg",
        network_fabric_name="example-fabric",
        body={
            "location": "eastuseuap",
            "properties": {
                "annotation": "annotation",
                "fabricASN": 29249,
                "ipv4Prefix": "10.18.0.0/19",
                "ipv6Prefix": "3FFE:FFFF:0:CD40::/59",
                "managementNetworkConfiguration": {
                    "infrastructureVpnConfiguration": {
                        "networkToNetworkInterconnectId": "/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkFabrics/example-fabric/networkToNetworkInterconnects/example-nni",
                        "optionAProperties": {
                            "bfdConfiguration": {"intervalInMilliSeconds": 300, "multiplier": 10},
                            "mtu": 1501,
                            "peerASN": 1235,
                            "primaryIpv4Prefix": "10.0.0.12/30",
                            "primaryIpv6Prefix": "4FFE:FFFF:0:CD30::a8/127",
                            "secondaryIpv4Prefix": "20.0.0.13/30",
                            "secondaryIpv6Prefix": "6FFE:FFFF:0:CD30::ac/127",
                            "vlanId": 3001,
                        },
                        "optionBProperties": {
                            "exportRouteTargets": ["65046:10050"],
                            "importRouteTargets": ["65046:10050"],
                            "routeTargets": {
                                "exportIpv4RouteTargets": ["65046:10039"],
                                "exportIpv6RouteTargets": ["65046:10039"],
                                "importIpv4RouteTargets": ["65046:10039"],
                                "importIpv6RouteTargets": ["65046:10039"],
                            },
                        },
                        "peeringOption": "OptionA",
                    },
                    "workloadVpnConfiguration": {
                        "networkToNetworkInterconnectId": "/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkFabrics/example-fabric/networkToNetworkInterconnects/example-nni",
                        "optionAProperties": {
                            "bfdConfiguration": {"intervalInMilliSeconds": 300, "multiplier": 5},
                            "mtu": 1500,
                            "peerASN": 61234,
                            "primaryIpv4Prefix": "10.0.0.14/30",
                            "primaryIpv6Prefix": "2FFE:FFFF:0:CD30::a7/126",
                            "secondaryIpv4Prefix": "10.0.0.15/30",
                            "secondaryIpv6Prefix": "2FFE:FFFF:0:CD30::ac/126",
                            "vlanId": 3000,
                        },
                        "optionBProperties": {
                            "exportRouteTargets": ["65046:10050"],
                            "importRouteTargets": ["65046:10050"],
                            "routeTargets": {
                                "exportIpv4RouteTargets": ["65046:10039"],
                                "exportIpv6RouteTargets": ["65046:10039"],
                                "importIpv4RouteTargets": ["65046:10039"],
                                "importIpv6RouteTargets": ["65046:10039"],
                            },
                        },
                        "peeringOption": "OptionA",
                    },
                },
                "networkFabricControllerId": "/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkFabricControllers/example-fabricController",
                "networkFabricSku": "M4-A400-A100-C16-aa",
                "rackCount": 4,
                "serverCountPerRack": 8,
                "terminalServerConfiguration": {
                    "password": "xxxx",
                    "primaryIpv4Prefix": "10.0.0.12/30",
                    "primaryIpv6Prefix": "4FFE:FFFF:0:CD30::a8/127",
                    "secondaryIpv4Prefix": "20.0.0.13/30",
                    "secondaryIpv6Prefix": "6FFE:FFFF:0:CD30::ac/127",
                    "serialNumber": "123456",
                    "username": "username",
                },
            },
            "tags": {"keyID": "keyValue"},
        },
    ).result()
    print(response)


# x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/NetworkFabrics_Create_MaximumSet_Gen.json
if __name__ == "__main__":
    main()
