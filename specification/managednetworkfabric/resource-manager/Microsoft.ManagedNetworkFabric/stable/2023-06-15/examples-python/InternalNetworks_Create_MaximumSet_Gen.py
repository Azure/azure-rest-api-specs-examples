from azure.identity import DefaultAzureCredential
from azure.mgmt.managednetworkfabric import ManagedNetworkFabricMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-managednetworkfabric
# USAGE
    python internal_networks_create_maximum_set_gen.py

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

    response = client.internal_networks.begin_create(
        resource_group_name="example-rg",
        l3_isolation_domain_name="example-l3domain",
        internal_network_name="example-internalnetwork",
        body={
            "properties": {
                "annotation": "annotation",
                "bgpConfiguration": {
                    "allowAS": 10,
                    "allowASOverride": "Enable",
                    "annotation": "annotation",
                    "bfdConfiguration": {"intervalInMilliSeconds": 300, "multiplier": 5},
                    "defaultRouteOriginate": "True",
                    "ipv4ListenRangePrefixes": ["10.1.0.0/25"],
                    "ipv4NeighborAddress": [{"address": "10.1.0.0"}],
                    "ipv6ListenRangePrefixes": ["2fff::/66"],
                    "ipv6NeighborAddress": [{"address": "2fff::"}],
                    "peerASN": 61234,
                },
                "connectedIPv4Subnets": [{"annotation": "annotation", "prefix": "10.0.0.0/24"}],
                "connectedIPv6Subnets": [{"annotation": "annotation", "prefix": "3FFE:FFFF:0:CD30::a0/29"}],
                "egressAclId": "/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/accessControlLists/example-acl",
                "exportRoutePolicy": {
                    "exportIpv4RoutePolicyId": "/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/routePolicies/routePolicyName",
                    "exportIpv6RoutePolicyId": "/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/routePolicies/routePolicyName",
                },
                "exportRoutePolicyId": "/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/routePolicies/routePolicyName",
                "extension": "NoExtension",
                "importRoutePolicy": {
                    "importIpv4RoutePolicyId": "/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/routePolicies/routePolicyName",
                    "importIpv6RoutePolicyId": "/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/routePolicies/routePolicyName",
                },
                "importRoutePolicyId": "/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/routePolicies/routePolicyName",
                "ingressAclId": "/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/accessControlLists/example-acl",
                "isMonitoringEnabled": "True",
                "mtu": 1500,
                "staticRouteConfiguration": {
                    "bfdConfiguration": {"intervalInMilliSeconds": 300, "multiplier": 15},
                    "extension": "NoExtension",
                    "ipv4Routes": [{"nextHop": ["10.0.0.1"], "prefix": "jffgck"}],
                    "ipv6Routes": [{"nextHop": ["3ffe::1"], "prefix": "2fff::/64"}],
                },
                "vlanId": 755,
            }
        },
    ).result()
    print(response)


# x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/InternalNetworks_Create_MaximumSet_Gen.json
if __name__ == "__main__":
    main()
