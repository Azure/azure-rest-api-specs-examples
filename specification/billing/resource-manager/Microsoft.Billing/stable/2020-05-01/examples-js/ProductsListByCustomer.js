const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the products for a customer. These don't include products billed based on usage.The operation is supported only for billing accounts with agreement type Microsoft Partner Agreement.
 *
 * @summary Lists the products for a customer. These don't include products billed based on usage.The operation is supported only for billing accounts with agreement type Microsoft Partner Agreement.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/ProductsListByCustomer.json
 */
async function productsListByInvoiceSection() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const billingAccountName = "{billingAccountName}";
  const customerName = "{customerName}";
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.products.listByCustomer(billingAccountName, customerName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

productsListByInvoiceSection().catch(console.error);
