const { MariaDBManagementClient } = require("@azure/arm-mariadb");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieve the Query-Store query texts for specified queryIds.
 *
 * @summary Retrieve the Query-Store query texts for specified queryIds.
 * x-ms-original-file: specification/mariadb/resource-manager/Microsoft.DBforMariaDB/stable/2018-06-01/examples/QueryTextsListByServer.json
 */
async function queryTextsListByServer() {
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const resourceGroupName = "testResourceGroupName";
  const serverName = "testServerName";
  const queryIds = ["1", "2"];
  const credential = new DefaultAzureCredential();
  const client = new MariaDBManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.queryTexts.listByServer(resourceGroupName, serverName, queryIds)) {
    resArray.push(item);
  }
  console.log(resArray);
}

queryTextsListByServer().catch(console.error);
