const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the invoices for a subscription. The operation is supported for billing accounts with agreement type Microsoft Partner Agreement or Microsoft Customer Agreement.
 *
 * @summary Lists the invoices for a subscription. The operation is supported for billing accounts with agreement type Microsoft Partner Agreement or Microsoft Customer Agreement.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/invoicesListByBillingSubscription.json
 */
async function invoicesListByBillingSubscription() {
  const subscriptionId =
    process.env["BILLING_SUBSCRIPTION_ID"] || "11111111-1111-1111-1111-111111111111";
  const periodStartDate = new Date("2023-01-01");
  const periodEndDate = new Date("2023-06-30");
  const options = {
    periodStartDate,
    periodEndDate,
  };
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.invoices.listByBillingSubscription(options)) {
    resArray.push(item);
  }
  console.log(resArray);
}
