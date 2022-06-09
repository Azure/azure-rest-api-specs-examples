```javascript
const { PeeringManagementClient } = require("@azure/arm-peering");
const { DefaultAzureCredential } = require("@azure/identity");

async function createADirectPeering() {
  const subscriptionId = "subId";
  const resourceGroupName = "rgName";
  const peeringName = "peeringName";
  const peering = {
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
      peerAsn: {
        id: "/subscriptions/subId/providers/Microsoft.Peering/peerAsns/myAsn1",
      },
    },
    kind: "Direct",
    location: "eastus",
    peeringLocation: "peeringLocation0",
    sku: { name: "Basic_Direct_Free" },
  };
  const credential = new DefaultAzureCredential();
  const client = new PeeringManagementClient(credential, subscriptionId);
  const result = await client.peerings.createOrUpdate(resourceGroupName, peeringName, peering);
  console.log(result);
}

createADirectPeering().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-peering_2.0.1/sdk/peering/arm-peering/README.md) on how to add the SDK to your project and authenticate.
