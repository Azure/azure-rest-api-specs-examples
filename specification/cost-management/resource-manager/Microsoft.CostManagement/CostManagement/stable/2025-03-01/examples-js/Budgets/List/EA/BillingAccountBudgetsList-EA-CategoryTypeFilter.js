const { CostManagementClient } = require("@azure/arm-costmanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to lists all budgets for the defined scope.
 *
 * @summary lists all budgets for the defined scope.
 * x-ms-original-file: 2025-03-01/Budgets/List/EA/BillingAccountBudgetsList-EA-CategoryTypeFilter.json
 */
async function billingAccountBudgetsListEACategoryTypeFilter() {
  const credential = new DefaultAzureCredential();
  const client = new CostManagementClient(credential);
  const resArray = new Array();
  for await (const item of client.budgets.list(
    "providers/Microsoft.Billing/billingAccounts/123456",
    { filter: "properties/category eq 'ReservationUtilization'" },
  )) {
    resArray.push(item);
  }

  console.log(resArray);
}
