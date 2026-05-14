const { PeeringManagementClient } = require("@azure/arm-peering");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates a new peering or updates an existing peering with the specified name under the given subscription and resource group.
 *
 * @summary creates a new peering or updates an existing peering with the specified name under the given subscription and resource group.
 * x-ms-original-file: 2025-05-01/CreateExchangePeering.json
 */
async function createAnExchangePeering() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "subId";
  const client = new PeeringManagementClient(credential, subscriptionId);
  const result = await client.peerings.createOrUpdate("rgName", "peeringName", {
    kind: "Exchange",
    location: "eastus",
    connectivityProbes: [{ azureRegion: "eastus", endpoint: "192.168.0.1", protocol: "ICMP" }],
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
      peerAsn: { id: "/subscriptions/subId/providers/Microsoft.Peering/peerAsns/myAsn1" },
    },
    peeringLocation: "peeringLocation0",
    sku: { name: "Basic_Exchange_Free" },
  });
  console.log(result);
}
