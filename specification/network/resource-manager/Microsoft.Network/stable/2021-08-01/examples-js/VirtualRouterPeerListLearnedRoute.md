Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-network_28.0.0/sdk/network/arm-network/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieves a list of routes the virtual hub bgp connection has learned.
 *
 * @summary Retrieves a list of routes the virtual hub bgp connection has learned.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/VirtualRouterPeerListLearnedRoute.json
 */
async function virtualRouterPeerListLearnedRoutes() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const hubName = "virtualRouter1";
  const connectionName = "peer1";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.virtualHubBgpConnections.beginListLearnedRoutesAndWait(
    resourceGroupName,
    hubName,
    connectionName
  );
  console.log(result);
}

virtualRouterPeerListLearnedRoutes().catch(console.error);
```
