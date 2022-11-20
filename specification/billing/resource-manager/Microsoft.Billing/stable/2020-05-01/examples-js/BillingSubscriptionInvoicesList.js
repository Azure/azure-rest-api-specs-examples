const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the invoices for a subscription.
 *
 * @summary Lists the invoices for a subscription.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/BillingSubscriptionInvoicesList.json
 */
async function billingSubscriptionsListByBillingAccount() {
  const subscriptionId = "{subscriptionId}";
  const periodStartDate = "2018-01-01";
  const periodEndDate = "2018-06-30";
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.invoices.listByBillingSubscription(
    periodStartDate,
    periodEndDate
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

billingSubscriptionsListByBillingAccount().catch(console.error);
