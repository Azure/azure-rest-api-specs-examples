const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a subscription by its ID. The operation is supported for billing accounts with agreement type Microsoft Customer Agreement and Microsoft Partner Agreement.
 *
 * @summary Gets a subscription by its ID. The operation is supported for billing accounts with agreement type Microsoft Customer Agreement and Microsoft Partner Agreement.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/BillingSubscription.json
 */
async function billingSubscription() {
  const subscriptionId = "{subscriptionId}";
  const billingAccountName = "{billingAccountName}";
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential, subscriptionId);
  const result = await client.billingSubscriptions.get(billingAccountName);
  console.log(result);
}

billingSubscription().catch(console.error);
