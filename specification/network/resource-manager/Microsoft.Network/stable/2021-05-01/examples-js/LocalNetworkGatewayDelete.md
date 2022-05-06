Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-network_27.0.0/sdk/network/arm-network/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the specified local network gateway.
 *
 * @summary Deletes the specified local network gateway.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/LocalNetworkGatewayDelete.json
 */
async function deleteLocalNetworkGateway() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const localNetworkGatewayName = "localgw";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.localNetworkGateways.beginDeleteAndWait(
    resourceGroupName,
    localNetworkGatewayName
  );
  console.log(result);
}

deleteLocalNetworkGateway().catch(console.error);
```
