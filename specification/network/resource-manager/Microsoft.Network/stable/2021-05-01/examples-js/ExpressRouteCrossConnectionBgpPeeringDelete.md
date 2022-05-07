Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-network_27.0.0/sdk/network/arm-network/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the specified peering from the ExpressRouteCrossConnection.
 *
 * @summary Deletes the specified peering from the ExpressRouteCrossConnection.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/ExpressRouteCrossConnectionBgpPeeringDelete.json
 */
async function deleteExpressRouteCrossConnectionBgpPeering() {
  const subscriptionId = "subid";
  const resourceGroupName = "CrossConnection-SiliconValley";
  const crossConnectionName = "<circuitServiceKey>";
  const peeringName = "AzurePrivatePeering";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.expressRouteCrossConnectionPeerings.beginDeleteAndWait(
    resourceGroupName,
    crossConnectionName,
    peeringName
  );
  console.log(result);
}

deleteExpressRouteCrossConnectionBgpPeering().catch(console.error);
```
