const { CostManagementClient } = require("@azure/arm-costmanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to gets the budget for the scope by budget name.
 *
 * @summary gets the budget for the scope by budget name.
 * x-ms-original-file: 2025-03-01/Budgets/Get/ReservationUtilization/Get-ReservationUtilization-AlertRule.json
 */
async function getReservationUtilizationAlertRule() {
  const credential = new DefaultAzureCredential();
  const client = new CostManagementClient(credential);
  const result = await client.budgets.get(
    "providers/Microsoft.Billing/billingAccounts/aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee:ffffffff-gggg-hhhh-iiii-jjjjjjjjjjjj_2023-04-01/billingProfiles/KKKK-LLLL-MMM-NNN",
    "TestAlertRule",
  );
  console.log(result);
}
