const { PeeringManagementClient } = require("@azure/arm-peering");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a new peering or updates an existing peering with the specified name under the given subscription and resource group.
 *
 * @summary Creates a new peering or updates an existing peering with the specified name under the given subscription and resource group.
 * x-ms-original-file: specification/peering/resource-manager/Microsoft.Peering/stable/2021-06-01/examples/CreateExchangePeering.json
 */
async function createAnExchangePeering() {
  const subscriptionId = "subId";
  const resourceGroupName = "rgName";
  const peeringName = "peeringName";
  const peering = {
    exchange: {
      connections: [
        {
          bgpSession: {
            maxPrefixesAdvertisedV4: 1000,
            maxPrefixesAdvertisedV6: 100,
            md5AuthenticationKey: "test-md5-auth-key",
            peerSessionIPv4Address: "192.168.2.1",
            peerSessionIPv6Address: "fd00::1",
          },
          connectionIdentifier: "CE495334-0E94-4E51-8164-8116D6CD284D",
          peeringDBFacilityId: 99999,
        },
        {
          bgpSession: {
            maxPrefixesAdvertisedV4: 1000,
            maxPrefixesAdvertisedV6: 100,
            md5AuthenticationKey: "test-md5-auth-key",
            peerSessionIPv4Address: "192.168.2.2",
            peerSessionIPv6Address: "fd00::2",
          },
          connectionIdentifier: "CDD8E673-CB07-47E6-84DE-3739F778762B",
          peeringDBFacilityId: 99999,
        },
      ],
      peerAsn: {
        id: "/subscriptions/subId/providers/Microsoft.Peering/peerAsns/myAsn1",
      },
    },
    kind: "Exchange",
    location: "eastus",
    peeringLocation: "peeringLocation0",
    sku: { name: "Basic_Exchange_Free" },
  };
  const credential = new DefaultAzureCredential();
  const client = new PeeringManagementClient(credential, subscriptionId);
  const result = await client.peerings.createOrUpdate(resourceGroupName, peeringName, peering);
  console.log(result);
}

createAnExchangePeering().catch(console.error);
