const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the billing profiles that a user has access to. The operation is supported for billing accounts with agreement type Microsoft Customer Agreement or Microsoft Partner Agreement.
 *
 * @summary Lists the billing profiles that a user has access to. The operation is supported for billing accounts with agreement type Microsoft Customer Agreement or Microsoft Partner Agreement.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/BillingProfilesListByBillingAccount.json
 */
async function billingProfilesListByBillingAccount() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const billingAccountName = "{billingAccountName}";
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.billingProfiles.listByBillingAccount(billingAccountName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

billingProfilesListByBillingAccount().catch(console.error);
