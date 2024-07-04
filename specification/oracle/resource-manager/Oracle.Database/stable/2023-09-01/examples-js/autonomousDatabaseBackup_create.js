const { OracleDatabaseManagementClient } = require("@azure/arm-oracledatabase");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create a AutonomousDatabaseBackup
 *
 * @summary Create a AutonomousDatabaseBackup
 * x-ms-original-file: specification/oracle/resource-manager/Oracle.Database/stable/2023-09-01/examples/autonomousDatabaseBackup_create.json
 */
async function createAutonomousDatabaseBackup() {
  const subscriptionId =
    process.env["ORACLEDATABASE_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["ORACLEDATABASE_RESOURCE_GROUP"] || "rg000";
  const autonomousdatabasename = "databasedb1";
  const adbbackupid = "1711644130";
  const resource = {
    properties: {
      autonomousDatabaseOcid: "ocid1.autonomousdatabase.oc1..aaaaa3klq",
      displayName: "Nightly Backup",
      retentionPeriodInDays: 365,
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new OracleDatabaseManagementClient(credential, subscriptionId);
  const result = await client.autonomousDatabaseBackups.beginCreateOrUpdateAndWait(
    resourceGroupName,
    autonomousdatabasename,
    adbbackupid,
    resource,
  );
  console.log(result);
}
