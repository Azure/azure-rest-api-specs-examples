const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the specified route from a route table.
 *
 * @summary Deletes the specified route from a route table.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/RouteTableRouteDelete.json
 */
async function deleteRoute() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const routeTableName = "testrt";
  const routeName = "route1";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.routes.beginDeleteAndWait(
    resourceGroupName,
    routeTableName,
    routeName
  );
  console.log(result);
}

deleteRoute().catch(console.error);
