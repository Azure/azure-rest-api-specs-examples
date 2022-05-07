Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-network_27.0.0/sdk/network/arm-network/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The VirtualNetworkGatewayConnectionResetSharedKey operation resets the virtual network gateway connection shared key for passed virtual network gateway connection in the specified resource group through Network resource provider.
 *
 * @summary The VirtualNetworkGatewayConnectionResetSharedKey operation resets the virtual network gateway connection shared key for passed virtual network gateway connection in the specified resource group through Network resource provider.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/VirtualNetworkGatewayConnectionResetSharedKey.json
 */
async function resetVirtualNetworkGatewayConnectionSharedKey() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const virtualNetworkGatewayConnectionName = "conn1";
  const parameters = { keyLength: 128 };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.virtualNetworkGatewayConnections.beginResetSharedKeyAndWait(
    resourceGroupName,
    virtualNetworkGatewayConnectionName,
    parameters
  );
  console.log(result);
}

resetVirtualNetworkGatewayConnectionSharedKey().catch(console.error);
```
