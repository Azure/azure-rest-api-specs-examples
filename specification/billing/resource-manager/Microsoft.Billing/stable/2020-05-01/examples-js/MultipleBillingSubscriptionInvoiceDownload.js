const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a URL to download multiple invoice documents (invoice pdf, tax receipts, credit notes) as a zip file.
 *
 * @summary Gets a URL to download multiple invoice documents (invoice pdf, tax receipts, credit notes) as a zip file.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/MultipleBillingSubscriptionInvoiceDownload.json
 */
async function billingSubscriptionInvoiceDownload() {
  const subscriptionId = "{subscriptionId}";
  const downloadUrls = [
    "https://management.azure.com/providers/Microsoft.Billing/billingAccounts/default/billingSubscriptions/{subscriptionId}/invoices/{invoiceName}/download?downloadToken={downloadToken}&useCache=True&api-version=2020-05-01",
    "https://management.azure.com/providers/Microsoft.Billing/billingAccounts/default/billingSubscriptions/{subscriptionId}/invoices/{invoiceName}/download?downloadToken={downloadToken}&useCache=True&api-version=2020-05-01",
    "https://management.azure.com/providers/Microsoft.Billing/billingAccounts/default/billingSubscriptions/{subscriptionId}/invoices/{invoiceName}/download?downloadToken={downloadToken}&useCache=True&api-version=2020-05-01",
  ];
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential, subscriptionId);
  const result = await client.invoices.beginDownloadMultipleBillingSubscriptionInvoicesAndWait(
    downloadUrls
  );
  console.log(result);
}

billingSubscriptionInvoiceDownload().catch(console.error);
