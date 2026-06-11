const { CostManagementClient } = require("@azure/arm-costmanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to lists all budgets for the defined scope.
 *
 * @summary lists all budgets for the defined scope.
 * x-ms-original-file: 2025-03-01/Budgets/List/MCA/BillingProfileBudgetsList-CategoryTypeFilter.json
 */
async function billingProfileBudgetsListMCACategoryTypeFilter() {
  const credential = new DefaultAzureCredential();
  const client = new CostManagementClient(credential);
  const resArray = new Array();
  for await (const item of client.budgets.list(
    "providers/Microsoft.Billing/billingAccounts/aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee:ffffffff-gggg-hhhh-iiii-jjjjjjjjjjjj_2023-04-01/billingProfiles/MYDEVTESTBP",
    { filter: "properties/category eq 'ReservationUtilization'" },
  )) {
    resArray.push(item);
  }

  console.log(resArray);
}
