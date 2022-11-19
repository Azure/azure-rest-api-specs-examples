const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the billing accounts that a user has access to.
 *
 * @summary Lists the billing accounts that a user has access to.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/BillingAccountsList.json
 */
async function billingAccountsList() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.billingAccounts.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

billingAccountsList().catch(console.error);
