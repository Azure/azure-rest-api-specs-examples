const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the transaction summary for an invoice. Transactions include purchases, refunds and Azure usage charges.
 *
 * @summary Gets the transaction summary for an invoice. Transactions include purchases, refunds and Azure usage charges.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/transactionSummaryGetByInvoice.json
 */
async function transactionSummaryGetByInvoice() {
  const billingAccountName =
    "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31";
  const invoiceName = "G123456789";
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential);
  const result = await client.transactions.getTransactionSummaryByInvoice(
    billingAccountName,
    invoiceName,
  );
  console.log(result);
}
