const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a enrollment account by name.
 *
 * @summary Gets a enrollment account by name.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/preview/2018-03-01-preview/examples/EnrollmentAccountsGet.json
 */
async function enrollmentAccountsGet() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const name = "e1bf1c8c-5ac6-44a0-bdcd-aa7c1cf60556";
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential, subscriptionId);
  const result = await client.enrollmentAccounts.get(name);
  console.log(result);
}

enrollmentAccountsGet().catch(console.error);
