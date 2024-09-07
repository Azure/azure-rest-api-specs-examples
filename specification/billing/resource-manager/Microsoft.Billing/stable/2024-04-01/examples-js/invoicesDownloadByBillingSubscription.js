const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a URL to download an invoice by billing subscription. The operation is supported for billing accounts with agreement type Microsoft Partner Agreement or Microsoft Customer Agreement.
 *
 * @summary Gets a URL to download an invoice by billing subscription. The operation is supported for billing accounts with agreement type Microsoft Partner Agreement or Microsoft Customer Agreement.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/invoicesDownloadByBillingSubscription.json
 */
async function invoicesDownloadByBillingSubscription() {
  const subscriptionId =
    process.env["BILLING_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const invoiceName = "E123456789";
  const documentName = "12345678";
  const options = {
    documentName,
  };
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential, subscriptionId);
  const result = await client.invoices.beginDownloadByBillingSubscriptionAndWait(
    invoiceName,
    options,
  );
  console.log(result);
}
