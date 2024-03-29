const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the subscriptions for a customer. The operation is supported only for billing accounts with agreement type Microsoft Partner Agreement.
 *
 * @summary Lists the subscriptions for a customer. The operation is supported only for billing accounts with agreement type Microsoft Partner Agreement.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/BillingSubscriptionsListByCustomer.json
 */
async function billingSubscriptionsListByCustomer() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const billingAccountName = "{billingAccountName}";
  const customerName = "{customerName}";
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.billingSubscriptions.listByCustomer(
    billingAccountName,
    customerName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

billingSubscriptionsListByCustomer().catch(console.error);
