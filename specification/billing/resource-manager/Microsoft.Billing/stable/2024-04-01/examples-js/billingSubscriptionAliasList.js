const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the subscription aliases for a billing account. The operation is supported for seat based billing subscriptions.
 *
 * @summary Lists the subscription aliases for a billing account. The operation is supported for seat based billing subscriptions.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/billingSubscriptionAliasList.json
 */
async function billingSubscriptionAliasList() {
  const billingAccountName =
    "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31";
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential);
  const resArray = new Array();
  for await (let item of client.billingSubscriptionsAliases.listByBillingAccount(
    billingAccountName,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
