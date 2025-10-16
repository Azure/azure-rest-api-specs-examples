const { OracleDatabaseManagementClient } = require("@azure/arm-oracledatabase");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to update a AutonomousDatabaseBackup
 *
 * @summary update a AutonomousDatabaseBackup
 * x-ms-original-file: 2025-09-01/autonomousDatabaseBackup_patch.json
 */
async function autonomousDatabaseBackupsUpdate() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new OracleDatabaseManagementClient(credential, subscriptionId);
  const result = await client.autonomousDatabaseBackups.update(
    "rg000",
    "databasedb1",
    "1711644130",
    {},
  );
  console.log(result);
}
