const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the subscriptions that are billed to an invoice section. The operation is supported only for billing accounts with agreement type Microsoft Customer Agreement.
 *
 * @summary Lists the subscriptions that are billed to an invoice section. The operation is supported only for billing accounts with agreement type Microsoft Customer Agreement.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/BillingSubscriptionsListByInvoiceSection.json
 */
async function billingSubscriptionsListByInvoiceSection() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const billingAccountName = "{billingAccountName}";
  const billingProfileName = "{billingProfileName}";
  const invoiceSectionName = "{invoiceSectionName}";
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.billingSubscriptions.listByInvoiceSection(
    billingAccountName,
    billingProfileName,
    invoiceSectionName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

billingSubscriptionsListByInvoiceSection().catch(console.error);
