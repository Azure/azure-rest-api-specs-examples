from azure.identity import DefaultAzureCredential
from azure.mgmt.managednetworkfabric import ManagedNetworkFabricMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-managednetworkfabric
# USAGE
    python route_policies_create_maximum_set_gen.py

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

    response = client.route_policies.begin_create(
        resource_group_name="example-rg",
        route_policy_name="example-routePolicy",
        body={
            "location": "eastus",
            "properties": {
                "addressFamilyType": "IPv4",
                "annotation": "annotation",
                "networkFabricId": "/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkFabrics/example-fabric",
                "statements": [
                    {
                        "action": {
                            "actionType": "Permit",
                            "ipCommunityProperties": {
                                "add": {
                                    "ipCommunityIds": [
                                        "/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/ipCommunities/example-ipCommunity"
                                    ]
                                },
                                "delete": {
                                    "ipCommunityIds": [
                                        "/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/ipCommunities/example-ipCommunity"
                                    ]
                                },
                                "set": {
                                    "ipCommunityIds": [
                                        "/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/ipCommunities/example-ipCommunity"
                                    ]
                                },
                            },
                            "ipExtendedCommunityProperties": {
                                "add": {
                                    "ipExtendedCommunityIds": [
                                        "/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/ipExtendedCommunities/example-ipExtendedCommunity"
                                    ]
                                },
                                "delete": {
                                    "ipExtendedCommunityIds": [
                                        "/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/ipExtendedCommunities/example-ipExtendedCommunity"
                                    ]
                                },
                                "set": {
                                    "ipExtendedCommunityIds": [
                                        "/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/ipExtendedCommunities/example-ipExtendedCommunity"
                                    ]
                                },
                            },
                            "localPreference": 20,
                        },
                        "annotation": "annotation",
                        "condition": {
                            "ipCommunityIds": [
                                "/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/ipCommunities/example-ipCommunity"
                            ],
                            "ipExtendedCommunityIds": [
                                "/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/ipExtendedCommunities/example-ipExtendedCommunity"
                            ],
                            "ipPrefixId": "/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/ipPrefixes/example-ipPrefix",
                            "type": "Or",
                        },
                        "sequenceNumber": 7,
                    }
                ],
            },
            "tags": {"keyID": "keyValue"},
        },
    ).result()
    print(response)


# x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/RoutePolicies_Create_MaximumSet_Gen.json
if __name__ == "__main__":
    main()
