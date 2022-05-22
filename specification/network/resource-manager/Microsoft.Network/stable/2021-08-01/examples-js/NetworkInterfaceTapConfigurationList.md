Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-network_28.0.0/sdk/network/arm-network/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get all Tap configurations in a network interface.
 *
 * @summary Get all Tap configurations in a network interface.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/NetworkInterfaceTapConfigurationList.json
 */
async function listVirtualNetworkTapConfigurations() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const networkInterfaceName = "mynic";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.networkInterfaceTapConfigurations.list(
    resourceGroupName,
    networkInterfaceName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

listVirtualNetworkTapConfigurations().catch(console.error);
```
