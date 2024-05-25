const { OracleDatabaseManagementClient } = require("@azure/arm-oracledatabase");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a AutonomousDatabaseNationalCharacterSet
 *
 * @summary Get a AutonomousDatabaseNationalCharacterSet
 * x-ms-original-file: specification/oracle/resource-manager/Oracle.Database/preview/2023-09-01-preview/examples/autonomousDatabaseNationalCharacterSet_get.json
 */
async function getAutonomousDbNationalCharacterSet() {
  const subscriptionId =
    process.env["ORACLEDATABASE_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const location = "eastus";
  const adbsncharsetname = "NATIONAL";
  const credential = new DefaultAzureCredential();
  const client = new OracleDatabaseManagementClient(credential, subscriptionId);
  const result = await client.autonomousDatabaseNationalCharacterSets.get(
    location,
    adbsncharsetname,
  );
  console.log(result);
}
