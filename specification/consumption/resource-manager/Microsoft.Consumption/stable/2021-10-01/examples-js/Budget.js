const { ConsumptionManagementClient } = require("@azure/arm-consumption");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the budget for the scope by budget name.
 *
 * @summary Gets the budget for the scope by budget name.
 * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/Budget.json
 */
async function budget() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const scope = "subscriptions/00000000-0000-0000-0000-000000000000";
  const budgetName = "TestBudget";
  const credential = new DefaultAzureCredential();
  const client = new ConsumptionManagementClient(credential, subscriptionId);
  const result = await client.budgets.get(scope, budgetName);
  console.log(result);
}

budget().catch(console.error);
