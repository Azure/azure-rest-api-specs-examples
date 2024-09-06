const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates the policies for a customer at billing account scope. This operation is supported only for billing accounts with agreement type Microsoft Partner Agreement.
 *
 * @summary Updates the policies for a customer at billing account scope. This operation is supported only for billing accounts with agreement type Microsoft Partner Agreement.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/policiesPutByCustomerAtBillingAccount.json
 */
async function policiesPutByCustomerAtBillingAccount() {
  const billingAccountName =
    "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31";
  const customerName = "11111111-1111-1111-1111-111111111111";
  const parameters = { properties: { viewCharges: "Allowed" } };
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential);
  const result = await client.policies.beginCreateOrUpdateByCustomerAtBillingAccountAndWait(
    billingAccountName,
    customerName,
    parameters,
  );
  console.log(result);
}
