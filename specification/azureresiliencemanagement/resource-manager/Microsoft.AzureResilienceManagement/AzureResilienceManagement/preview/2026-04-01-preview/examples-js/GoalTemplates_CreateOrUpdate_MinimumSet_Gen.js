const { AzureResilienceManagementClient } = require("@azure/arm-resiliencemanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create a GoalTemplate
 *
 * @summary create a GoalTemplate
 * x-ms-original-file: 2026-04-01-preview/GoalTemplates_CreateOrUpdate_MinimumSet_Gen.json
 */
async function goalTemplatesCreateOrUpdateMinimumSet() {
  const credential = new DefaultAzureCredential();
  const client = new AzureResilienceManagementClient(credential);
  const result = await client.goalTemplates.createOrUpdate("sg1", "gt1", {
    properties: { goalType: "Resiliency" },
  });
  console.log(result);
}
