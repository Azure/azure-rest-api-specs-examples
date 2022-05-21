Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-network_28.0.0/sdk/network/arm-network/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the specified virtual network by resource group.
 *
 * @summary Gets the specified virtual network by resource group.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/VirtualNetworkGetWithSubnetDelegation.json
 */
async function getVirtualNetworkWithADelegatedSubnet() {
  const subscriptionId = "subId";
  const resourceGroupName = "rg1";
  const virtualNetworkName = "test-vnet";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.virtualNetworks.get(resourceGroupName, virtualNetworkName);
  console.log(result);
}

getVirtualNetworkWithADelegatedSubnet().catch(console.error);
```
