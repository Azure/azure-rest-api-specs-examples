const { OracleDatabaseManagementClient } = require("@azure/arm-oracledatabase");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Generate wallet action on Autonomous Database
 *
 * @summary Generate wallet action on Autonomous Database
 * x-ms-original-file: specification/oracle/resource-manager/Oracle.Database/preview/2023-09-01-preview/examples/autonomousDatabase_generateWallet.json
 */
async function generateWalletActionOnAutonomousDatabase() {
  const subscriptionId =
    process.env["ORACLEDATABASE_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["ORACLEDATABASE_RESOURCE_GROUP"] || "rg000";
  const autonomousdatabasename = "databasedb1";
  const body = {
    generateType: "Single",
    isRegional: false,
    password: "ExamplePassword12345",
  };
  const credential = new DefaultAzureCredential();
  const client = new OracleDatabaseManagementClient(credential, subscriptionId);
  const result = await client.autonomousDatabases.generateWallet(
    resourceGroupName,
    autonomousdatabasename,
    body,
  );
  console.log(result);
}
