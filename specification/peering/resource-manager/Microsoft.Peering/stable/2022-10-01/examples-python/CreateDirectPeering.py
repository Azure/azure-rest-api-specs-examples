from azure.identity import DefaultAzureCredential
from azure.mgmt.peering import PeeringManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-peering
# USAGE
    python create_a_direct_peering.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = PeeringManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="subId",
    )

    response = client.peerings.create_or_update(
        resource_group_name="rgName",
        peering_name="peeringName",
        peering={
            "kind": "Direct",
            "location": "eastus",
            "properties": {
                "direct": {
                    "connections": [
                        {
                            "bandwidthInMbps": 10000,
                            "bgpSession": {
                                "maxPrefixesAdvertisedV4": 1000,
                                "maxPrefixesAdvertisedV6": 100,
                                "md5AuthenticationKey": "test-md5-auth-key",
                                "sessionPrefixV4": "192.168.0.0/31",
                                "sessionPrefixV6": "fd00::0/127",
                            },
                            "connectionIdentifier": "5F4CB5C7-6B43-4444-9338-9ABC72606C16",
                            "peeringDBFacilityId": 99999,
                            "sessionAddressProvider": "Peer",
                            "useForPeeringService": False,
                        },
                        {
                            "bandwidthInMbps": 10000,
                            "connectionIdentifier": "8AB00818-D533-4504-A25A-03A17F61201C",
                            "peeringDBFacilityId": 99999,
                            "sessionAddressProvider": "Microsoft",
                            "useForPeeringService": True,
                        },
                    ],
                    "directPeeringType": "Edge",
                    "peerAsn": {"id": "/subscriptions/subId/providers/Microsoft.Peering/peerAsns/myAsn1"},
                },
                "peeringLocation": "peeringLocation0",
            },
            "sku": {"name": "Basic_Direct_Free"},
        },
    )
    print(response)


# x-ms-original-file: specification/peering/resource-manager/Microsoft.Peering/stable/2022-10-01/examples/CreateDirectPeering.json
if __name__ == "__main__":
    main()
