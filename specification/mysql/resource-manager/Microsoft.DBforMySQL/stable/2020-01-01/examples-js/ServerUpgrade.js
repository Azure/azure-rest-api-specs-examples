const { MySQLManagementClient } = require("@azure/arm-mysql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Upgrade server version.
 *
 * @summary Upgrade server version.
 * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2020-01-01/examples/ServerUpgrade.json
 */
async function serverUpgrade() {
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const resourceGroupName = "TestGroup";
  const serverName = "testserver";
  const parameters = { targetServerVersion: "5.7" };
  const credential = new DefaultAzureCredential();
  const client = new MySQLManagementClient(credential, subscriptionId);
  const result = await client.servers.beginUpgradeAndWait(
    resourceGroupName,
    serverName,
    parameters
  );
  console.log(result);
}

serverUpgrade().catch(console.error);
