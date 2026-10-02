const { AzureResilienceManagementClient } = require("@azure/arm-resiliencemanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to get a GoalResource
 *
 * @summary get a GoalResource
 * x-ms-original-file: 2026-08-31-preview/GoalResources_Get_MaximumSet_Gen.json
 */
async function goalResourcesGetMaximumSet() {
  const credential = new DefaultAzureCredential();
  const client = new AzureResilienceManagementClient(credential);
  const result = await client.goalResources.get("umyghwnfpzsgrhpczizcn", "ga1", "gr1");
  console.log(result);
}
