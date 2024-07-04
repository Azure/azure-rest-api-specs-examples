const { OracleDatabaseManagementClient } = require("@azure/arm-oracledatabase");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a AutonomousDatabaseCharacterSet
 *
 * @summary Get a AutonomousDatabaseCharacterSet
 * x-ms-original-file: specification/oracle/resource-manager/Oracle.Database/stable/2023-09-01/examples/autonomousDatabaseCharacterSet_get.json
 */
async function getAutonomousDbCharacterSet() {
  const subscriptionId =
    process.env["ORACLEDATABASE_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const location = "eastus";
  const adbscharsetname = "DATABASE";
  const credential = new DefaultAzureCredential();
  const client = new OracleDatabaseManagementClient(credential, subscriptionId);
  const result = await client.autonomousDatabaseCharacterSets.get(location, adbscharsetname);
  console.log(result);
}
