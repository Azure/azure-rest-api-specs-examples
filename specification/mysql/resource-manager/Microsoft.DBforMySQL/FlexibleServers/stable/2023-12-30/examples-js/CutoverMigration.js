const { MySQLManagementFlexibleServerClient } = require("@azure/arm-mysql-flexible");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Cutover migration for MySQL import, it will switch source elastic server DNS to flexible server.
 *
 * @summary Cutover migration for MySQL import, it will switch source elastic server DNS to flexible server.
 * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/FlexibleServers/stable/2023-12-30/examples/CutoverMigration.json
 */
async function cutoverMigrationForMySqlImport() {
  const subscriptionId =
    process.env["MYSQL_SUBSCRIPTION_ID"] || "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const resourceGroupName = process.env["MYSQL_RESOURCE_GROUP"] || "testrg";
  const serverName = "mysqltestserver";
  const credential = new DefaultAzureCredential();
  const client = new MySQLManagementFlexibleServerClient(credential, subscriptionId);
  const result = await client.serversMigration.beginCutoverMigrationAndWait(
    resourceGroupName,
    serverName,
  );
  console.log(result);
}
