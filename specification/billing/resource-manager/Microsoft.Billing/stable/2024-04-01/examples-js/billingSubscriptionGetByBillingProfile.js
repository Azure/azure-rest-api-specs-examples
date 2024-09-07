const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a subscription by its billing profile and ID. The operation is supported for billing accounts with agreement type Enterprise Agreement.
 *
 * @summary Gets a subscription by its billing profile and ID. The operation is supported for billing accounts with agreement type Enterprise Agreement.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/billingSubscriptionGetByBillingProfile.json
 */
async function billingSubscriptionGetByBillingProfile() {
  const billingAccountName = "pcn.94077792";
  const billingProfileName = "6478903";
  const billingSubscriptionName = "6b96d3f2-9008-4a9d-912f-f87744185aa3";
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential);
  const result = await client.billingSubscriptions.getByBillingProfile(
    billingAccountName,
    billingProfileName,
    billingSubscriptionName,
  );
  console.log(result);
}
