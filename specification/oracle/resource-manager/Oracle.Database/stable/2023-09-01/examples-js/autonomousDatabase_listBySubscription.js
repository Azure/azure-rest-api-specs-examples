const { OracleDatabaseManagementClient } = require("@azure/arm-oracledatabase");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List AutonomousDatabase resources by subscription ID
 *
 * @summary List AutonomousDatabase resources by subscription ID
 * x-ms-original-file: specification/oracle/resource-manager/Oracle.Database/stable/2023-09-01/examples/autonomousDatabase_listBySubscription.json
 */
async function listAutonomousDatabaseBySubscription() {
  const subscriptionId =
    process.env["ORACLEDATABASE_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new OracleDatabaseManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.autonomousDatabases.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}
