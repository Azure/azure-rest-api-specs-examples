const { MariaDBManagementClient } = require("@azure/arm-mariadb");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Starts a stopped server.
 *
 * @summary Starts a stopped server.
 * x-ms-original-file: specification/mariadb/resource-manager/Microsoft.DBforMariaDB/stable/2020-01-01/examples/ServerStart.json
 */
async function serverStart() {
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const resourceGroupName = "TestGroup";
  const serverName = "testserver";
  const credential = new DefaultAzureCredential();
  const client = new MariaDBManagementClient(credential, subscriptionId);
  const result = await client.servers.beginStartAndWait(resourceGroupName, serverName);
  console.log(result);
}

serverStart().catch(console.error);
