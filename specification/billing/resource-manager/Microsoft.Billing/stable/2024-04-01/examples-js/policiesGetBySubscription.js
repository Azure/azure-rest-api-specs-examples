const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the policies that are managed by the Billing Admin for the defined subscriptions. This is supported for Microsoft Online Services Program, Microsoft Customer Agreement and Microsoft Partner Agreement.
 *
 * @summary Lists the policies that are managed by the Billing Admin for the defined subscriptions. This is supported for Microsoft Online Services Program, Microsoft Customer Agreement and Microsoft Partner Agreement.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/policiesGetBySubscription.json
 */
async function policiesGetBySubscription() {
  const subscriptionId =
    process.env["BILLING_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential, subscriptionId);
  const result = await client.policies.getBySubscription();
  console.log(result);
}
