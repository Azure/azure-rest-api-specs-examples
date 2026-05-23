const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to gets a billing account by its ID.
 *
 * @summary gets a billing account by its ID.
 * x-ms-original-file: 2024-04-01/billingAccountsGetWithExpand.json
 */
async function billingAccountsGetWithExpand() {
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential);
  const result = await client.billingAccounts.get(
    "10000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31",
  );
  console.log(result);
}
