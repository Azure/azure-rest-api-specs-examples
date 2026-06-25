const { AzureResilienceManagementClient } = require("@azure/arm-resiliencemanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to get a GoalResource
 *
 * @summary get a GoalResource
 * x-ms-original-file: 2026-04-01-preview/GoalResources_Get_MinimumSet_Gen.json
 */
async function goalResourcesGetMinimumSet() {
  const credential = new DefaultAzureCredential();
  const client = new AzureResilienceManagementClient(credential);
  const result = await client.goalResources.get("sg1", "ga1", "gr1");
  console.log(result);
}
