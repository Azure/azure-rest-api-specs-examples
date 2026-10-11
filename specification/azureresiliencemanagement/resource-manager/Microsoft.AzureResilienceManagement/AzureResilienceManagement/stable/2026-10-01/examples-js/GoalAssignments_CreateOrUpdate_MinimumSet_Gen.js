const { AzureResilienceManagementClient } = require("@azure/arm-resiliencemanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates a goal assignment.
 *
 * @summary creates or updates a goal assignment.
 * x-ms-original-file: 2026-10-01/GoalAssignments_CreateOrUpdate_MinimumSet_Gen.json
 */
async function goalAssignmentsCreateOrUpdateMinimumSet() {
  const credential = new DefaultAzureCredential();
  const client = new AzureResilienceManagementClient(credential);
  await client.goalAssignments.createOrUpdate("production-sg", "zonal-resiliency-goal", {
    properties: { requireZonalResiliency: true },
  });
}
