const { OracleDatabaseManagementClient } = require("@azure/arm-oracledatabase");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to list AutonomousDatabase resources by resource group
 *
 * @summary list AutonomousDatabase resources by resource group
 * x-ms-original-file: 2025-09-01/AutonomousDatabases_ListByResourceGroup_MaximumSet_Gen.json
 */
async function listAutonomousDatabaseByResourceGroupGeneratedByMaximumSetRule() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new OracleDatabaseManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.autonomousDatabases.listByResourceGroup("rgopenapi")) {
    resArray.push(item);
  }

  console.log(resArray);
}
