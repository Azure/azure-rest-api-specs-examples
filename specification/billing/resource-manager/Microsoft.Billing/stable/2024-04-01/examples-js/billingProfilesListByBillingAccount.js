const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the billing profiles that a user has access to. The operation is supported for billing accounts with agreement of type Microsoft Customer Agreement and Microsoft Partner Agreement.
 *
 * @summary Lists the billing profiles that a user has access to. The operation is supported for billing accounts with agreement of type Microsoft Customer Agreement and Microsoft Partner Agreement.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/billingProfilesListByBillingAccount.json
 */
async function billingProfilesListByBillingAccount() {
  const billingAccountName =
    "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31";
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential);
  const resArray = new Array();
  for await (let item of client.billingProfiles.listByBillingAccount(billingAccountName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
