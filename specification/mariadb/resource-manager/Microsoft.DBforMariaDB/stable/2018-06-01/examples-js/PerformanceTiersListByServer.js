const { MariaDBManagementClient } = require("@azure/arm-mariadb");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List all the performance tiers for a MariaDB server.
 *
 * @summary List all the performance tiers for a MariaDB server.
 * x-ms-original-file: specification/mariadb/resource-manager/Microsoft.DBforMariaDB/stable/2018-06-01/examples/PerformanceTiersListByServer.json
 */
async function performanceTiersList() {
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const resourceGroupName = "testrg";
  const serverName = "mariadbtestsvc1";
  const credential = new DefaultAzureCredential();
  const client = new MariaDBManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.serverBasedPerformanceTier.list(resourceGroupName, serverName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

performanceTiersList().catch(console.error);
