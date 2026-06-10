const { CostManagementClient } = require("@azure/arm-costmanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to delete cost allocation rule for billing account or enterprise enrollment.
 *
 * @summary delete cost allocation rule for billing account or enterprise enrollment.
 * x-ms-original-file: 2025-03-01/CostAllocationRuleDelete.json
 */
async function deleteCostAllocationRule() {
  const credential = new DefaultAzureCredential();
  const client = new CostManagementClient(credential);
  await client.costAllocationRules.delete("100", "testRule");
}
