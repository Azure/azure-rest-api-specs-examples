const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to gets a billing account by its ID.
 *
 * @summary gets a billing account by its ID.
 * x-ms-original-file: 2024-04-01/billingAccountWithExpandForPONumber.json
 */
async function billingAccountWithExpandForPONumber() {
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential);
  const result = await client.billingAccounts.get("8608480");
  console.log(result);
}
