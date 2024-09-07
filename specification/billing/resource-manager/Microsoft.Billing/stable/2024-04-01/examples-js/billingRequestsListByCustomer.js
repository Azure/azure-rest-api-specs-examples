const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The list of billing requests submitted for the customer.
 *
 * @summary The list of billing requests submitted for the customer.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/billingRequestsListByCustomer.json
 */
async function billingRequestsListByCustomer() {
  const billingAccountName =
    "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31";
  const billingProfileName = "xxxx-xxxx-xxx-xxx";
  const customerName = "11111111-1111-1111-1111-111111111111";
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential);
  const resArray = new Array();
  for await (let item of client.billingRequests.listByCustomer(
    billingAccountName,
    billingProfileName,
    customerName,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
