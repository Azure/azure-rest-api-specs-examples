Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-network_27.0.0/sdk/network/arm-network/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieves the details of a RouteTable.
 *
 * @summary Retrieves the details of a RouteTable.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/HubRouteTableGet.json
 */
async function routeTableGet() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const virtualHubName = "virtualHub1";
  const routeTableName = "hubRouteTable1";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.hubRouteTables.get(resourceGroupName, virtualHubName, routeTableName);
  console.log(result);
}

routeTableGet().catch(console.error);
```
