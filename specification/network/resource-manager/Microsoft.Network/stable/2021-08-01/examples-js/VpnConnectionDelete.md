Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-network_28.0.0/sdk/network/arm-network/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a vpn connection.
 *
 * @summary Deletes a vpn connection.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/VpnConnectionDelete.json
 */
async function vpnConnectionDelete() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const gatewayName = "gateway1";
  const connectionName = "vpnConnection1";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.vpnConnections.beginDeleteAndWait(
    resourceGroupName,
    gatewayName,
    connectionName
  );
  console.log(result);
}

vpnConnectionDelete().catch(console.error);
```
