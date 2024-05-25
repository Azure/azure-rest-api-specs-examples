const { OracleDatabaseManagementClient } = require("@azure/arm-oracledatabase");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List AutonomousDatabaseBackup resources by AutonomousDatabase
 *
 * @summary List AutonomousDatabaseBackup resources by AutonomousDatabase
 * x-ms-original-file: specification/oracle/resource-manager/Oracle.Database/preview/2023-09-01-preview/examples/autonomousDatabaseBackup_listByParent.json
 */
async function listAutonomousDatabaseBackupsByAutonomousDatabase() {
  const subscriptionId =
    process.env["ORACLEDATABASE_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["ORACLEDATABASE_RESOURCE_GROUP"] || "rg000";
  const autonomousdatabasename = "databasedb1";
  const credential = new DefaultAzureCredential();
  const client = new OracleDatabaseManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.autonomousDatabaseBackups.listByAutonomousDatabase(
    resourceGroupName,
    autonomousdatabasename,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
