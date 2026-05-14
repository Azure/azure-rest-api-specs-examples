const { PeeringManagementClient } = require("@azure/arm-peering");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates a new peering or updates an existing peering with the specified name under the given subscription and resource group.
 *
 * @summary creates a new peering or updates an existing peering with the specified name under the given subscription and resource group.
 * x-ms-original-file: 2025-05-01/CreateDirectPeering.json
 */
async function createADirectPeering() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "subId";
  const client = new PeeringManagementClient(credential, subscriptionId);
  const result = await client.peerings.createOrUpdate("rgName", "peeringName", {
    kind: "Direct",
    location: "eastus",
    connectivityProbes: [{ azureRegion: "eastus", endpoint: "192.168.0.1", protocol: "ICMP" }],
    direct: {
      connections: [
        {
          bandwidthInMbps: 10000,
          bgpSession: {
            maxPrefixesAdvertisedV4: 1000,
            maxPrefixesAdvertisedV6: 100,
            md5AuthenticationKey: "test-md5-auth-key",
            sessionPrefixV4: "192.168.0.0/31",
            sessionPrefixV6: "fd00::0/127",
          },
          connectionIdentifier: "5F4CB5C7-6B43-4444-9338-9ABC72606C16",
          peeringDBFacilityId: 99999,
          sessionAddressProvider: "Peer",
          useForPeeringService: false,
        },
        {
          bandwidthInMbps: 10000,
          connectionIdentifier: "8AB00818-D533-4504-A25A-03A17F61201C",
          peeringDBFacilityId: 99999,
          sessionAddressProvider: "Microsoft",
          useForPeeringService: true,
        },
      ],
      directPeeringType: "Edge",
      peerAsn: { id: "/subscriptions/subId/providers/Microsoft.Peering/peerAsns/myAsn1" },
    },
    peeringLocation: "peeringLocation0",
    sku: { name: "Basic_Direct_Free" },
  });
  console.log(result);
}
