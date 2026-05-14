const { PeeringManagementClient } = require("@azure/arm-peering");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates a new peering or updates an existing peering with the specified name under the given subscription and resource group.
 *
 * @summary creates a new peering or updates an existing peering with the specified name under the given subscription and resource group.
 * x-ms-original-file: 2025-05-01/CreatePeeringWithExchangeRouteServer.json
 */
async function createAPeeringWithExchangeRouteServer() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "subId";
  const client = new PeeringManagementClient(credential, subscriptionId);
  const result = await client.peerings.createOrUpdate("rgName", "peeringName", {
    kind: "Direct",
    location: "eastus",
    connectivityProbes: [{ azureRegion: "eastus", endpoint: "192.168.0.1", protocol: "TCP" }],
    direct: {
      connections: [
        {
          bandwidthInMbps: 10000,
          bgpSession: {
            maxPrefixesAdvertisedV4: 1000,
            maxPrefixesAdvertisedV6: 100,
            microsoftSessionIPv4Address: "192.168.0.123",
            peerSessionIPv4Address: "192.168.0.234",
            sessionPrefixV4: "192.168.0.0/24",
          },
          connectionIdentifier: "5F4CB5C7-6B43-4444-9338-9ABC72606C16",
          peeringDBFacilityId: 99999,
          sessionAddressProvider: "Peer",
          useForPeeringService: true,
        },
      ],
      directPeeringType: "IxRs",
      peerAsn: { id: "/subscriptions/subId/providers/Microsoft.Peering/peerAsns/myAsn1" },
    },
    peeringLocation: "peeringLocation0",
    sku: { name: "Premium_Direct_Free" },
  });
  console.log(result);
}
