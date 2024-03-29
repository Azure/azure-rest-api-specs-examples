from azure.identity import DefaultAzureCredential
from azure.mgmt.managednetworkfabric import ManagedNetworkFabricMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-managednetworkfabric
# USAGE
    python external_networks_update_maximum_set_gen.py

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

    response = client.external_networks.begin_update(
        resource_group_name="example-rg",
        l3_isolation_domain_name="example-l3domain",
        external_network_name="example-externalnetwork",
        body={
            "properties": {
                "annotation": "annotation1",
                "exportRoutePolicy": {
                    "exportIpv4RoutePolicyId": "/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/routePolicies/routePolicyName",
                    "exportIpv6RoutePolicyId": "/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/routePolicies/routePolicyName",
                },
                "exportRoutePolicyId": "/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/routePolicies/routePolicyName",
                "importRoutePolicy": {
                    "importIpv4RoutePolicyId": "/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/routePolicies/routePolicyName",
                    "importIpv6RoutePolicyId": "/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/routePolicies/routePolicyName",
                },
                "importRoutePolicyId": "/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/routePolicies/routePolicyName",
                "optionAProperties": {
                    "bfdConfiguration": {"intervalInMilliSeconds": 300, "multiplier": 15},
                    "egressAclId": "/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/accessControlLists/example-acl",
                    "ingressAclId": "/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/accessControlLists/example-acl",
                    "mtu": 1500,
                    "peerASN": 65047,
                    "primaryIpv4Prefix": "10.1.1.0/30",
                    "primaryIpv6Prefix": "3FFE:FFFF:0:CD30::a0/126",
                    "secondaryIpv4Prefix": "10.1.1.4/30",
                    "secondaryIpv6Prefix": "3FFE:FFFF:0:CD30::a4/126",
                    "vlanId": 1001,
                },
                "optionBProperties": {
                    "exportRouteTargets": ["65046:10039"],
                    "importRouteTargets": ["65046:10039"],
                    "routeTargets": {
                        "exportIpv4RouteTargets": ["65046:10039"],
                        "exportIpv6RouteTargets": ["65046:10039"],
                        "importIpv4RouteTargets": ["65046:10039"],
                        "importIpv6RouteTargets": ["65046:10039"],
                    },
                },
                "peeringOption": "OptionA",
            }
        },
    ).result()
    print(response)


# x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/ExternalNetworks_Update_MaximumSet_Gen.json
if __name__ == "__main__":
    main()
