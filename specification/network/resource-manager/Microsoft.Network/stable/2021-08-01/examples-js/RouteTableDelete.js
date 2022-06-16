const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the specified route table.
 *
 * @summary Deletes the specified route table.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/RouteTableDelete.json
 */
async function deleteRouteTable() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const routeTableName = "testrt";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.routeTables.beginDeleteAndWait(resourceGroupName, routeTableName);
  console.log(result);
}

deleteRouteTable().catch(console.error);
