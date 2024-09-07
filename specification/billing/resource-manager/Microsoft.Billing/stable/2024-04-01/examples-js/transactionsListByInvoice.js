const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the transactions for an invoice. Transactions include purchases, refunds and Azure usage charges.
 *
 * @summary Lists the transactions for an invoice. Transactions include purchases, refunds and Azure usage charges.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/transactionsListByInvoice.json
 */
async function transactionsListByInvoice() {
  const billingAccountName =
    "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31";
  const invoiceName = "G123456789";
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential);
  const resArray = new Array();
  for await (let item of client.transactions.listByInvoice(billingAccountName, invoiceName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
