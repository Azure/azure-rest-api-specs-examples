const { OracleDatabaseManagementClient } = require("@azure/arm-oracledatabase");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to perform switchover action on Autonomous Database
 *
 * @summary perform switchover action on Autonomous Database
 * x-ms-original-file: 2025-09-01/AutonomousDatabases_Switchover_MaximumSet_Gen.json
 */
async function performSwitchoverActionOnAutonomousDatabaseGeneratedByMaximumSetRule() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new OracleDatabaseManagementClient(credential, subscriptionId);
  const result = await client.autonomousDatabases.switchover("rgopenapi", "databasedb1", {
    peerDbId: "peerDbId",
    peerDbOcid: "yozpqyefqhirkybmzwgoidyl",
    peerDbLocation: "cxlzbzbfzi",
  });
  console.log(result);
}
