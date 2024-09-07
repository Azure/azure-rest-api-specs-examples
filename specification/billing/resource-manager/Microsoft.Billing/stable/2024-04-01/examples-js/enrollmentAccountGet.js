const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets an enrollment account by ID. The operation is supported only for billing accounts with agreement type Enterprise Agreement.
 *
 * @summary Gets an enrollment account by ID. The operation is supported only for billing accounts with agreement type Enterprise Agreement.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/enrollmentAccountGet.json
 */
async function enrollmentAccountGet() {
  const billingAccountName = "6564892";
  const enrollmentAccountName = "257698";
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential);
  const result = await client.enrollmentAccounts.get(billingAccountName, enrollmentAccountName);
  console.log(result);
}
