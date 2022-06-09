```javascript
const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a peering in the specified ExpressRouteCrossConnection.
 *
 * @summary Creates or updates a peering in the specified ExpressRouteCrossConnection.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/ExpressRouteCrossConnectionBgpPeeringCreate.json
 */
async function expressRouteCrossConnectionBgpPeeringCreate() {
  const subscriptionId = "subid";
  const resourceGroupName = "CrossConnection-SiliconValley";
  const crossConnectionName = "<circuitServiceKey>";
  const peeringName = "AzurePrivatePeering";
  const peeringParameters = {
    ipv6PeeringConfig: {
      primaryPeerAddressPrefix: "3FFE:FFFF:0:CD30::/126",
      secondaryPeerAddressPrefix: "3FFE:FFFF:0:CD30::4/126",
    },
    peerASN: 200,
    primaryPeerAddressPrefix: "192.168.16.252/30",
    secondaryPeerAddressPrefix: "192.168.18.252/30",
    vlanId: 200,
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.expressRouteCrossConnectionPeerings.beginCreateOrUpdateAndWait(
    resourceGroupName,
    crossConnectionName,
    peeringName,
    peeringParameters
  );
  console.log(result);
}

expressRouteCrossConnectionBgpPeeringCreate().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-network_28.0.0/sdk/network/arm-network/README.md) on how to add the SDK to your project and authenticate.
