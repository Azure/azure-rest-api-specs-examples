const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a savings plan order by billing account.
 *
 * @summary Get a savings plan order by billing account.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/savingsPlanOrderExpandedScheduleGetByBillingAccount.json
 */
async function savingsPlanOrderWithExpandedPaymentsGet() {
  const billingAccountName =
    "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31";
  const savingsPlanOrderId = "20000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential);
  const result = await client.savingsPlanOrders.getByBillingAccount(
    billingAccountName,
    savingsPlanOrderId,
  );
  console.log(result);
}
