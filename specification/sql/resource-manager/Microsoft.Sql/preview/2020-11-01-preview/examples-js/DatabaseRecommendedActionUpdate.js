const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates a database recommended action.
 *
 * @summary Updates a database recommended action.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/DatabaseRecommendedActionUpdate.json
 */
async function updateDatabaseRecommendedAction() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "workloadinsight-demos";
  const serverName = "misosisvr";
  const databaseName = "IndexAdvisor_test_3";
  const advisorName = "CreateIndex";
  const recommendedActionName = "IR_[CRM]_[DataPoints]_4821CD2F9510D98184BB";
  const parameters = { state: { currentValue: "Pending" } };
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.databaseRecommendedActions.update(
    resourceGroupName,
    serverName,
    databaseName,
    advisorName,
    recommendedActionName,
    parameters
  );
  console.log(result);
}

updateDatabaseRecommendedAction().catch(console.error);
