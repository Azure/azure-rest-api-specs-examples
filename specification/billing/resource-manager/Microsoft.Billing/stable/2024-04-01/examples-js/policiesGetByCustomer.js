const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the policies for a customer. This operation is supported only for billing accounts with agreement type Microsoft Partner Agreement.
 *
 * @summary Lists the policies for a customer. This operation is supported only for billing accounts with agreement type Microsoft Partner Agreement.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/policiesGetByCustomer.json
 */
async function policiesGetByCustomer() {
  const billingAccountName =
    "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31";
  const billingProfileName = "xxxx-xxxx-xxx-xxx";
  const customerName = "11111111-1111-1111-1111-111111111111";
  const policyName = "default";
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential);
  const result = await client.policies.getByCustomer(
    billingAccountName,
    billingProfileName,
    customerName,
    policyName,
  );
  console.log(result);
}
