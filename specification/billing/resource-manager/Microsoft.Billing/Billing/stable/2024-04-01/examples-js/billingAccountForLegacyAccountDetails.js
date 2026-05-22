const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to lists the billing accounts that a user has access to.
 *
 * @summary lists the billing accounts that a user has access to.
 * x-ms-original-file: 2024-04-01/billingAccountForLegacyAccountDetails.json
 */
async function billingAccountForLegacyAccountDetails() {
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential);
  const resArray = new Array();
  for await (const item of client.billingAccounts.list()) {
    resArray.push(item);
  }

  console.log(resArray);
}
