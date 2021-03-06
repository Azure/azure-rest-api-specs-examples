const { MySQLManagementClient } = require("@azure/arm-mysql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List all the performance tiers for a MySQL server.
 *
 * @summary List all the performance tiers for a MySQL server.
 * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2017-12-01/examples/PerformanceTiersListByServer.json
 */
async function performanceTiersList() {
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const resourceGroupName = "testrg";
  const serverName = "mysqltestsvc1";
  const credential = new DefaultAzureCredential();
  const client = new MySQLManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.serverBasedPerformanceTier.list(resourceGroupName, serverName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

performanceTiersList().catch(console.error);
