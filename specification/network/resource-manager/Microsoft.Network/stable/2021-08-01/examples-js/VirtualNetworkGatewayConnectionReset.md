Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-network_28.0.0/sdk/network/arm-network/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Resets the virtual network gateway connection specified.
 *
 * @summary Resets the virtual network gateway connection specified.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/VirtualNetworkGatewayConnectionReset.json
 */
async function resetVirtualNetworkGatewayConnection() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const virtualNetworkGatewayConnectionName = "conn1";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.virtualNetworkGatewayConnections.beginResetConnectionAndWait(
    resourceGroupName,
    virtualNetworkGatewayConnectionName
  );
  console.log(result);
}

resetVirtualNetworkGatewayConnection().catch(console.error);
```
