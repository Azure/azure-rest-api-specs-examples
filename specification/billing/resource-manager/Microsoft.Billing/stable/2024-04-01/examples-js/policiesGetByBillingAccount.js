const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the policies for a billing account of Enterprise Agreement type.
 *
 * @summary Get the policies for a billing account of Enterprise Agreement type.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/policiesGetByBillingAccount.json
 */
async function policiesGetByBillingAccount() {
  const billingAccountName = "1234567";
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential);
  const result = await client.policies.getByBillingAccount(billingAccountName);
  console.log(result);
}
