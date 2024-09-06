const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates the billing property of a subscription. Currently, cost center can be updated for billing accounts with agreement type Microsoft Customer Agreement and subscription service usage address can be updated for billing accounts with agreement type Microsoft Online Service Program.
 *
 * @summary Updates the billing property of a subscription. Currently, cost center can be updated for billing accounts with agreement type Microsoft Customer Agreement and subscription service usage address can be updated for billing accounts with agreement type Microsoft Online Service Program.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/billingPropertyPatchCostCenter.json
 */
async function billingPropertyPatchCostCenter() {
  const subscriptionId =
    process.env["BILLING_SUBSCRIPTION_ID"] || "11111111-1111-1111-1111-111111111111";
  const parameters = { properties: { costCenter: "1010" } };
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential, subscriptionId);
  const result = await client.billingPropertyOperations.update(parameters);
  console.log(result);
}
