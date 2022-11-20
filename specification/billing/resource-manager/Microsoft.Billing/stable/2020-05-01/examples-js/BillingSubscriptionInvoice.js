const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets an invoice by subscription ID and invoice ID.
 *
 * @summary Gets an invoice by subscription ID and invoice ID.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/BillingSubscriptionInvoice.json
 */
async function billingSubscriptionsListByBillingAccount() {
  const subscriptionId = "{subscriptionId}";
  const invoiceName = "{invoiceName}";
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential, subscriptionId);
  const result = await client.invoices.getBySubscriptionAndInvoiceId(invoiceName);
  console.log(result);
}

billingSubscriptionsListByBillingAccount().catch(console.error);
