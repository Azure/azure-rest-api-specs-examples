const { OracleDatabaseManagementClient } = require("@azure/arm-oracledatabase");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to perform ChangeDisasterRecoveryConfiguration action on Autonomous Database
 *
 * @summary perform ChangeDisasterRecoveryConfiguration action on Autonomous Database
 * x-ms-original-file: 2025-09-01/AutonomousDatabases_ChangeDisasterRecoveryConfiguration_MaximumSet_Gen.json
 */
async function performChangeDisasterRecoveryConfigurationActionOnAutonomousDatabaseGeneratedByMaximumSetRule() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new OracleDatabaseManagementClient(credential, subscriptionId);
  const result = await client.autonomousDatabases.changeDisasterRecoveryConfiguration(
    "rgopenapi",
    "databasedb1",
    {
      disasterRecoveryType: "Adg",
      isReplicateAutomaticBackups: true,
      timeSnapshotStandbyEnabledTill: new Date("2025-08-01T04:32:58.725Z"),
      isSnapshotStandby: true,
    },
  );
  console.log(result);
}
