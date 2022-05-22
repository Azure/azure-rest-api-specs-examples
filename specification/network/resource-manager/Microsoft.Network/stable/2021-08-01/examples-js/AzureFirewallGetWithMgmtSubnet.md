Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-network_28.0.0/sdk/network/arm-network/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the specified Azure Firewall.
 *
 * @summary Gets the specified Azure Firewall.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/AzureFirewallGetWithMgmtSubnet.json
 */
async function getAzureFirewallWithManagementSubnet() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const azureFirewallName = "azurefirewall";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.azureFirewalls.get(resourceGroupName, azureFirewallName);
  console.log(result);
}

getAzureFirewallWithManagementSubnet().catch(console.error);
```
