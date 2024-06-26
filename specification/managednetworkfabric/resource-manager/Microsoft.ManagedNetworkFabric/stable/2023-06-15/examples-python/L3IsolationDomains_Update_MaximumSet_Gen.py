from azure.identity import DefaultAzureCredential
from azure.mgmt.managednetworkfabric import ManagedNetworkFabricMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-managednetworkfabric
# USAGE
    python l3_isolation_domains_update_maximum_set_gen.py

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

    response = client.l3_isolation_domains.begin_update(
        resource_group_name="example-rg",
        l3_isolation_domain_name="example-l3domain",
        body={
            "properties": {
                "aggregateRouteConfiguration": {
                    "ipv4Routes": [{"prefix": "10.0.0.0/24"}],
                    "ipv6Routes": [{"prefix": "3FFE:FFFF:0:CD30::a0/29"}],
                },
                "annotation": "annotation1",
                "connectedSubnetRoutePolicy": {
                    "exportRoutePolicy": {
                        "exportIpv4RoutePolicyId": "/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/routePolicies/example-routePolicy1",
                        "exportIpv6RoutePolicyId": "/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/routePolicies/example-routePolicy1",
                    },
                    "exportRoutePolicyId": "/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/routePolicies/routePolicyName",
                },
                "redistributeConnectedSubnets": "True",
                "redistributeStaticRoutes": "False",
            },
            "tags": {"key4953": "1234"},
        },
    ).result()
    print(response)


# x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/L3IsolationDomains_Update_MaximumSet_Gen.json
if __name__ == "__main__":
    main()
