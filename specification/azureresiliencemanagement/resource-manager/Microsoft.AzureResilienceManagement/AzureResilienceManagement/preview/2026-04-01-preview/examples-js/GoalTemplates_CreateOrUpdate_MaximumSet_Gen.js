const { AzureResilienceManagementClient } = require("@azure/arm-resiliencemanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create a GoalTemplate
 *
 * @summary create a GoalTemplate
 * x-ms-original-file: 2026-04-01-preview/GoalTemplates_CreateOrUpdate_MaximumSet_Gen.json
 */
async function goalTemplatesCreateOrUpdateMaximumSet() {
  const credential = new DefaultAzureCredential();
  const client = new AzureResilienceManagementClient(credential);
  const result = await client.goalTemplates.createOrUpdate("zumt", "gt1", {
    properties: {
      requireHighAvailability: "Required",
      requireDisasterRecovery: "NotRequired",
      regionalRecoveryPointObjective: "PT15M",
      regionalRecoveryTimeObjective: "PT30M",
      goalType: "Resiliency",
    },
  });
  console.log(result);
}
