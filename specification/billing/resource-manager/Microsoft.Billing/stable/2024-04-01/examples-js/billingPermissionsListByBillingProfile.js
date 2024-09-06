const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the billing permissions the caller has on a billing profile.
 *
 * @summary Lists the billing permissions the caller has on a billing profile.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/billingPermissionsListByBillingProfile.json
 */
async function billingPermissionsListByBillingProfile() {
  const billingAccountName =
    "10000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31";
  const billingProfileName = "xxxx-xxxx-xxx-xxx";
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential);
  const resArray = new Array();
  for await (let item of client.billingPermissions.listByBillingProfile(
    billingAccountName,
    billingProfileName,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
