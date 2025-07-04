const { OracleDatabaseManagementClient } = require("@azure/arm-oracledatabase");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to perform switchover action on Autonomous Database
 *
 * @summary perform switchover action on Autonomous Database
 * x-ms-original-file: 2025-03-01/autonomousDatabase_switchover.json
 */
async function autonomousDatabasesSwitchover() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new OracleDatabaseManagementClient(credential, subscriptionId);
  const result = await client.autonomousDatabases.switchover("rg000", "databasedb1", {
    peerDbId: "peerDbId",
  });
  console.log(result);
}
