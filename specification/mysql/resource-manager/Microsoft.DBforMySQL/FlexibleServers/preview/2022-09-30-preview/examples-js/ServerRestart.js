const { MySQLManagementFlexibleServerClient } = require("@azure/arm-mysql-flexible");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Restarts a server.
 *
 * @summary Restarts a server.
 * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/FlexibleServers/preview/2022-09-30-preview/examples/ServerRestart.json
 */
async function restartAServer() {
  const subscriptionId =
    process.env["MYSQL_SUBSCRIPTION_ID"] || "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const resourceGroupName = process.env["MYSQL_RESOURCE_GROUP"] || "TestGroup";
  const serverName = "testserver";
  const parameters = {
    maxFailoverSeconds: 60,
    restartWithFailover: "Enabled",
  };
  const credential = new DefaultAzureCredential();
  const client = new MySQLManagementFlexibleServerClient(credential, subscriptionId);
  const result = await client.servers.beginRestartAndWait(
    resourceGroupName,
    serverName,
    parameters
  );
  console.log(result);
}
