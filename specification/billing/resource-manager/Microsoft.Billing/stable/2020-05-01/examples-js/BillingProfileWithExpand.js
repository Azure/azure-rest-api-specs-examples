const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a billing profile by its ID. The operation is supported for billing accounts with agreement type Microsoft Customer Agreement or Microsoft Partner Agreement.
 *
 * @summary Gets a billing profile by its ID. The operation is supported for billing accounts with agreement type Microsoft Customer Agreement or Microsoft Partner Agreement.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/BillingProfileWithExpand.json
 */
async function billingProfileWithExpand() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const billingAccountName = "{billingAccountName}";
  const billingProfileName = "{billingProfileName}";
  const expand = "invoiceSections";
  const options = { expand };
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential, subscriptionId);
  const result = await client.billingProfiles.get(billingAccountName, billingProfileName, options);
  console.log(result);
}

billingProfileWithExpand().catch(console.error);
