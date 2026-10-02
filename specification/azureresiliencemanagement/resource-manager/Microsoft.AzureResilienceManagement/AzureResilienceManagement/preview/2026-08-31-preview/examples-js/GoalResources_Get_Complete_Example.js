const { AzureResilienceManagementClient } = require("@azure/arm-resiliencemanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to get a GoalResource
 *
 * @summary get a GoalResource
 * x-ms-original-file: 2026-08-31-preview/GoalResources_Get_Complete_Example.json
 */
async function goalResourcesGetCompleteExample() {
  const credential = new DefaultAzureCredential();
  const client = new AzureResilienceManagementClient(credential);
  const result = await client.goalResources.get(
    "production-sg",
    "resiliencyGoalAssignment",
    "web-app-resource",
  );
  console.log(result);
}
