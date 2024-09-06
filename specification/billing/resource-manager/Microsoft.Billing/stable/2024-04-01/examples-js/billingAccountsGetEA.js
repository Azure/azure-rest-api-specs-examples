const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a billing account by its ID.
 *
 * @summary Gets a billing account by its ID.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/billingAccountsGetEA.json
 */
async function billingAccountsGetEa() {
  const billingAccountName = "6575495";
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential);
  const result = await client.billingAccounts.get(billingAccountName);
  console.log(result);
}
