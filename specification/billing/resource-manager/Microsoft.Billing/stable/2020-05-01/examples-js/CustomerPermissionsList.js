const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the billing permissions the caller has for a customer.
 *
 * @summary Lists the billing permissions the caller has for a customer.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/CustomerPermissionsList.json
 */
async function billingProfilePermissionsList() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const billingAccountName = "{billingAccountName}";
  const customerName = "{customerName}";
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.billingPermissions.listByCustomer(
    billingAccountName,
    customerName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

billingProfilePermissionsList().catch(console.error);
