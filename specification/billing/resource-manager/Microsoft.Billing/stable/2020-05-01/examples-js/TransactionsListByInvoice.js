const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the transactions for an invoice. Transactions include purchases, refunds and Azure usage charges.
 *
 * @summary Lists the transactions for an invoice. Transactions include purchases, refunds and Azure usage charges.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/TransactionsListByInvoice.json
 */
async function transactionsListByInvoice() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const billingAccountName = "{billingAccountName}";
  const invoiceName = "{invoiceName}";
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.transactions.listByInvoice(billingAccountName, invoiceName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

transactionsListByInvoice().catch(console.error);
