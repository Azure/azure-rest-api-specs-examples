const { MySQLManagementClient } = require("@azure/arm-mysql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets information about a server.
 *
 * @summary Gets information about a server.
 * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2017-12-01/examples/ServerGet.json
 */
async function serverGet() {
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const resourceGroupName = "testrg";
  const serverName = "mysqltestsvc4";
  const credential = new DefaultAzureCredential();
  const client = new MySQLManagementClient(credential, subscriptionId);
  const result = await client.servers.get(resourceGroupName, serverName);
  console.log(result);
}

serverGet().catch(console.error);
