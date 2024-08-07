const { OracleDatabaseManagementClient } = require("@azure/arm-oracledatabase");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update a AutonomousDatabaseBackup
 *
 * @summary Update a AutonomousDatabaseBackup
 * x-ms-original-file: specification/oracle/resource-manager/Oracle.Database/stable/2023-09-01/examples/autonomousDatabaseBackup_patch.json
 */
async function patchAutonomousDatabaseBackup() {
  const subscriptionId =
    process.env["ORACLEDATABASE_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["ORACLEDATABASE_RESOURCE_GROUP"] || "rg000";
  const autonomousdatabasename = "databasedb1";
  const adbbackupid = "1711644130";
  const properties = {};
  const credential = new DefaultAzureCredential();
  const client = new OracleDatabaseManagementClient(credential, subscriptionId);
  const result = await client.autonomousDatabaseBackups.beginUpdateAndWait(
    resourceGroupName,
    autonomousdatabasename,
    adbbackupid,
    properties,
  );
  console.log(result);
}
