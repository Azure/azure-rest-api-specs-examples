from azure.identity import DefaultAzureCredential
from azure.mgmt.peering import PeeringManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-peering
# USAGE
    python create_an_exchange_peering.py

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
            "kind": "Exchange",
            "location": "eastus",
            "properties": {
                "exchange": {
                    "connections": [
                        {
                            "bgpSession": {
                                "maxPrefixesAdvertisedV4": 1000,
                                "maxPrefixesAdvertisedV6": 100,
                                "md5AuthenticationKey": "test-md5-auth-key",
                                "peerSessionIPv4Address": "192.168.2.1",
                                "peerSessionIPv6Address": "fd00::1",
                            },
                            "connectionIdentifier": "CE495334-0E94-4E51-8164-8116D6CD284D",
                            "peeringDBFacilityId": 99999,
                        },
                        {
                            "bgpSession": {
                                "maxPrefixesAdvertisedV4": 1000,
                                "maxPrefixesAdvertisedV6": 100,
                                "md5AuthenticationKey": "test-md5-auth-key",
                                "peerSessionIPv4Address": "192.168.2.2",
                                "peerSessionIPv6Address": "fd00::2",
                            },
                            "connectionIdentifier": "CDD8E673-CB07-47E6-84DE-3739F778762B",
                            "peeringDBFacilityId": 99999,
                        },
                    ],
                    "peerAsn": {"id": "/subscriptions/subId/providers/Microsoft.Peering/peerAsns/myAsn1"},
                },
                "peeringLocation": "peeringLocation0",
            },
            "sku": {"name": "Basic_Exchange_Free"},
        },
    )
    print(response)


# x-ms-original-file: specification/peering/resource-manager/Microsoft.Peering/stable/2022-10-01/examples/CreateExchangePeering.json
if __name__ == "__main__":
    main()
