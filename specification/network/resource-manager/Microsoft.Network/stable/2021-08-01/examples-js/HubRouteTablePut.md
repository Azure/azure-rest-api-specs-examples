Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-network_28.0.0/sdk/network/arm-network/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a RouteTable resource if it doesn't exist else updates the existing RouteTable.
 *
 * @summary Creates a RouteTable resource if it doesn't exist else updates the existing RouteTable.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/HubRouteTablePut.json
 */
async function routeTablePut() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const virtualHubName = "virtualHub1";
  const routeTableName = "hubRouteTable1";
  const routeTableParameters = {
    labels: ["label1", "label2"],
    routes: [
      {
        name: "route1",
        destinationType: "CIDR",
        destinations: ["10.0.0.0/8", "20.0.0.0/8", "30.0.0.0/8"],
        nextHop:
          "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/azureFirewalls/azureFirewall1",
        nextHopType: "ResourceId",
      },
    ],
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.hubRouteTables.beginCreateOrUpdateAndWait(
    resourceGroupName,
    virtualHubName,
    routeTableName,
    routeTableParameters
  );
  console.log(result);
}

routeTablePut().catch(console.error);
```
