Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-network_28.0.0/sdk/network/arm-network/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists ExpressRoute gateways in a given resource group.
 *
 * @summary Lists ExpressRoute gateways in a given resource group.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/ExpressRouteGatewayListByResourceGroup.json
 */
async function expressRouteGatewayListByResourceGroup() {
  const subscriptionId = "subid";
  const resourceGroupName = "resourceGroupName";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.expressRouteGateways.listByResourceGroup(resourceGroupName);
  console.log(result);
}

expressRouteGatewayListByResourceGroup().catch(console.error);
```
