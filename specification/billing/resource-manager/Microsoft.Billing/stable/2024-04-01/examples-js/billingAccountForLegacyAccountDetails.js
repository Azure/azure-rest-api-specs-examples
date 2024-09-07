const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the billing accounts that a user has access to.
 *
 * @summary Lists the billing accounts that a user has access to.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/billingAccountForLegacyAccountDetails.json
 */
async function billingAccountForLegacyAccountDetails() {
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential);
  const resArray = new Array();
  for await (let item of client.billingAccounts.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}
