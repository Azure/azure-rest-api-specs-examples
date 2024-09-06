const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the policies for a customer at billing account scope. This operation is supported only for billing accounts with agreement type Microsoft Partner Agreement.
 *
 * @summary Lists the policies for a customer at billing account scope. This operation is supported only for billing accounts with agreement type Microsoft Partner Agreement.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/policiesGetByCustomerAtBillingAccount.json
 */
async function policiesGetByCustomerAtBillingAccount() {
  const billingAccountName =
    "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31";
  const customerName = "11111111-1111-1111-1111-111111111111";
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential);
  const result = await client.policies.getByCustomerAtBillingAccount(
    billingAccountName,
    customerName,
  );
  console.log(result);
}
