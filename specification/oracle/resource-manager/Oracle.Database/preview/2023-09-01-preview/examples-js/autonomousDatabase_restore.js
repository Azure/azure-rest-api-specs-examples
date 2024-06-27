const { OracleDatabaseManagementClient } = require("@azure/arm-oracledatabase");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Restores an Autonomous Database based on the provided request parameters.
 *
 * @summary Restores an Autonomous Database based on the provided request parameters.
 * x-ms-original-file: specification/oracle/resource-manager/Oracle.Database/preview/2023-09-01-preview/examples/autonomousDatabase_restore.json
 */
async function performRestoreActionOnAutonomousDatabase() {
  const subscriptionId =
    process.env["ORACLEDATABASE_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["ORACLEDATABASE_RESOURCE_GROUP"] || "rg000";
  const autonomousdatabasename = "databasedb1";
  const body = {
    timestamp: new Date("2024-04-23T00:00:00.000Z"),
  };
  const credential = new DefaultAzureCredential();
  const client = new OracleDatabaseManagementClient(credential, subscriptionId);
  const result = await client.autonomousDatabases.beginRestoreAndWait(
    resourceGroupName,
    autonomousdatabasename,
    body,
  );
  console.log(result);
}
