const { CostManagementClient } = require("@azure/arm-costmanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to lists all budgets for the defined scope.
 *
 * @summary lists all budgets for the defined scope.
 * x-ms-original-file: 2025-03-01/Budgets/List/RBAC/SubscriptionBudgetsList.json
 */
async function subscriptionBudgetsList() {
  const credential = new DefaultAzureCredential();
  const client = new CostManagementClient(credential);
  const resArray = new Array();
  for await (const item of client.budgets.list(
    "subscriptions/00000000-0000-0000-0000-000000000000",
  )) {
    resArray.push(item);
  }

  console.log(resArray);
}
