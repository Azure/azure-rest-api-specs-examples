const { OracleDatabaseManagementClient } = require("@azure/arm-oracledatabase");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to This operation shrinks the current allocated storage down to the current actual used data storage.
 *
 * @summary This operation shrinks the current allocated storage down to the current actual used data storage.
 * x-ms-original-file: specification/oracle/resource-manager/Oracle.Database/preview/2023-09-01-preview/examples/autonomousDatabase_shrink.json
 */
async function performShrinkActionOnAutonomousDatabase() {
  const subscriptionId =
    process.env["ORACLEDATABASE_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["ORACLEDATABASE_RESOURCE_GROUP"] || "rg000";
  const autonomousdatabasename = "databasedb1";
  const credential = new DefaultAzureCredential();
  const client = new OracleDatabaseManagementClient(credential, subscriptionId);
  const result = await client.autonomousDatabases.beginShrinkAndWait(
    resourceGroupName,
    autonomousdatabasename,
  );
  console.log(result);
}
