const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a URL to download an invoice.
 *
 * @summary Gets a URL to download an invoice.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/BillingSubscriptionInvoiceDownload.json
 */
async function billingSubscriptionInvoiceDownload() {
  const subscriptionId = "{subscriptionId}";
  const invoiceName = "{invoiceName}";
  const downloadToken = "DRS_12345";
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential, subscriptionId);
  const result = await client.invoices.beginDownloadBillingSubscriptionInvoiceAndWait(
    invoiceName,
    downloadToken
  );
  console.log(result);
}

billingSubscriptionInvoiceDownload().catch(console.error);
