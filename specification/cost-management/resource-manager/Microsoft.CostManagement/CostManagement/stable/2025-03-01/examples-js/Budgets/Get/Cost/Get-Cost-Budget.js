const { CostManagementClient } = require("@azure/arm-costmanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to gets the budget for the scope by budget name.
 *
 * @summary gets the budget for the scope by budget name.
 * x-ms-original-file: 2025-03-01/Budgets/Get/Cost/Get-Cost-Budget.json
 */
async function getCostBudget() {
  const credential = new DefaultAzureCredential();
  const client = new CostManagementClient(credential);
  const result = await client.budgets.get(
    "subscriptions/00000000-0000-0000-0000-000000000000",
    "TestBudget",
  );
  console.log(result);
}
