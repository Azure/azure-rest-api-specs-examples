const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The list of billing requests submitted by a user.
 *
 * @summary The list of billing requests submitted by a user.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/billingRequestsListByUser.json
 */
async function billingRequestsListByUser() {
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential);
  const resArray = new Array();
  for await (let item of client.billingRequests.listByUser()) {
    resArray.push(item);
  }
  console.log(resArray);
}
