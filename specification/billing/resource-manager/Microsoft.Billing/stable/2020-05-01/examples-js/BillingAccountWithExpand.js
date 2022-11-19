const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a billing account by its ID.
 *
 * @summary Gets a billing account by its ID.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/BillingAccountWithExpand.json
 */
async function billingAccountWithExpand() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const billingAccountName = "{billingAccountName}";
  const expand = "soldTo,billingProfiles,billingProfiles/invoiceSections";
  const options = { expand };
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential, subscriptionId);
  const result = await client.billingAccounts.get(billingAccountName, options);
  console.log(result);
}

billingAccountWithExpand().catch(console.error);
