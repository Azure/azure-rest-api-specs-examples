const { AzureResilienceManagementClient } = require("@azure/arm-resiliencemanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to deletes a goal template. Deprecated: the GoalTemplate resource type is deprecated and is removed in 2026-09-30-preview. Set resiliency intent directly on the GoalAssignment (requireZonalResiliency, requireRegionalResiliency) and read resource posture from GoalResource (zonalResiliency, regionalResiliency) instead.
 *
 * @summary deletes a goal template. Deprecated: the GoalTemplate resource type is deprecated and is removed in 2026-09-30-preview. Set resiliency intent directly on the GoalAssignment (requireZonalResiliency, requireRegionalResiliency) and read resource posture from GoalResource (zonalResiliency, regionalResiliency) instead.
 * x-ms-original-file: 2026-08-31-preview/GoalTemplates_Delete_MaximumSet_Gen.json
 */
async function goalTemplatesDeleteMaximumSet() {
  const credential = new DefaultAzureCredential();
  const client = new AzureResilienceManagementClient(credential);
  await client.goalTemplates.delete("ajsvdpsdgp", "gt1");
}
