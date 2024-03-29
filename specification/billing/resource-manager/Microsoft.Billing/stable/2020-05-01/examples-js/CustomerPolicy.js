const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the policies for a customer. This operation is supported only for billing accounts with agreement type Microsoft Partner Agreement.
 *
 * @summary Lists the policies for a customer. This operation is supported only for billing accounts with agreement type Microsoft Partner Agreement.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/CustomerPolicy.json
 */
async function policyByCustomer() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const billingAccountName = "{billingAccountName}";
  const customerName = "{customerName}";
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential, subscriptionId);
  const result = await client.policies.getByCustomer(billingAccountName, customerName);
  console.log(result);
}

policyByCustomer().catch(console.error);
