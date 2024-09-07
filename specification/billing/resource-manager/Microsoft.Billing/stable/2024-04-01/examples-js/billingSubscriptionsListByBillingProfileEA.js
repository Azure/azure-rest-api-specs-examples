const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the subscriptions that are billed to a billing profile. The operation is supported for billing accounts with agreement type Microsoft Customer Agreement or Microsoft Partner Agreement.
 *
 * @summary Lists the subscriptions that are billed to a billing profile. The operation is supported for billing accounts with agreement type Microsoft Customer Agreement or Microsoft Partner Agreement.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/billingSubscriptionsListByBillingProfileEA.json
 */
async function billingSubscriptionsListByBillingProfileEa() {
  const billingAccountName = "pcn.94077792";
  const billingProfileName = "6478903";
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential);
  const resArray = new Array();
  for await (let item of client.billingSubscriptions.listByBillingProfile(
    billingAccountName,
    billingProfileName,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
