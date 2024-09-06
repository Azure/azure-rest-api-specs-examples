const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the billing permissions the caller has for a customer at billing account level.
 *
 * @summary Lists the billing permissions the caller has for a customer at billing account level.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/billingPermissionsListByCustomerAtBillingAccount.json
 */
async function billingPermissionsListByCustomerAtBillingAccount() {
  const billingAccountName =
    "10000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31";
  const customerName = "11111111-1111-1111-1111-111111111111";
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential);
  const resArray = new Array();
  for await (let item of client.billingPermissions.listByCustomerAtBillingAccount(
    billingAccountName,
    customerName,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
