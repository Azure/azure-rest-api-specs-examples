const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the billing properties for a subscription
 *
 * @summary Gets the billing properties for a subscription
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/billingPropertyGetMPA.json
 */
async function billingPropertyGetMpa() {
  const subscriptionId =
    process.env["BILLING_SUBSCRIPTION_ID"] || "11111111-1111-1111-1111-111111111111";
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential, subscriptionId);
  const result = await client.billingPropertyOperations.get();
  console.log(result);
}
