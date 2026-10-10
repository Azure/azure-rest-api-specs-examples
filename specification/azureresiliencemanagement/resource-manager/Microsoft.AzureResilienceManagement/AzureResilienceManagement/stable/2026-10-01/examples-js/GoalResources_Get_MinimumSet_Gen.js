const { AzureResilienceManagementClient } = require("@azure/arm-resiliencemanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to gets a goal resource.
 *
 * @summary gets a goal resource.
 * x-ms-original-file: 2026-10-01/GoalResources_Get_MinimumSet_Gen.json
 */
async function goalResourcesGetMinimumSet() {
  const credential = new DefaultAzureCredential();
  const client = new AzureResilienceManagementClient(credential);
  const result = await client.goalResources.get(
    "production-sg",
    "zonal-resiliency-goal",
    "primary-vm",
  );
  console.log(result);
}
