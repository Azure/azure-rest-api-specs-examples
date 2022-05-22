Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-network_28.0.0/sdk/network/arm-network/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates the specified Virtual Router.
 *
 * @summary Creates or updates the specified Virtual Router.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/VirtualRouterPut.json
 */
async function createVirtualRouter() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const virtualRouterName = "virtualRouter";
  const parameters = {
    hostedGateway: {
      id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/vnetGateway",
    },
    location: "West US",
    tags: { key1: "value1" },
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.virtualRouters.beginCreateOrUpdateAndWait(
    resourceGroupName,
    virtualRouterName,
    parameters
  );
  console.log(result);
}

createVirtualRouter().catch(console.error);
```
