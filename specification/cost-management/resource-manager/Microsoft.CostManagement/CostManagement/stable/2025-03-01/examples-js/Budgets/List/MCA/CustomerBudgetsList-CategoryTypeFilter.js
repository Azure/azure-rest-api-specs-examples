const { CostManagementClient } = require("@azure/arm-costmanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to lists all budgets for the defined scope.
 *
 * @summary lists all budgets for the defined scope.
 * x-ms-original-file: 2025-03-01/Budgets/List/MCA/CustomerBudgetsList-CategoryTypeFilter.json
 */
async function customerBudgetsListMCACSPCategoryTypeFilter() {
  const credential = new DefaultAzureCredential();
  const client = new CostManagementClient(credential);
  const resArray = new Array();
  for await (const item of client.budgets.list(
    "providers/Microsoft.Billing/billingAccounts/aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee:ffffffff-gggg-hhhh-iiii-jjjjjjjjjjjj_2023-04-01/customers/000000-1111-2222-3333-444444444444",
    { filter: "properties/category eq 'ReservationUtilization'" },
  )) {
    resArray.push(item);
  }

  console.log(resArray);
}
