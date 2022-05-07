Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-network_27.0.0/sdk/network/arm-network/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieves the details of a virtual wan vpn gateway.
 *
 * @summary Retrieves the details of a virtual wan vpn gateway.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/VpnGatewayGet.json
 */
async function vpnGatewayGet() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const gatewayName = "gateway1";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.vpnGateways.get(resourceGroupName, gatewayName);
  console.log(result);
}

vpnGatewayGet().catch(console.error);
```
