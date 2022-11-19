const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a customer by its ID. The operation is supported only for billing accounts with agreement type Microsoft Partner Agreement.
 *
 * @summary Gets a customer by its ID. The operation is supported only for billing accounts with agreement type Microsoft Partner Agreement.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/CustomerWithExpand.json
 */
async function customerWithExpand() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const billingAccountName = "{billingAccountName}";
  const customerName = "{customerName}";
  const expand = "enabledAzurePlans,resellers";
  const options = { expand };
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential, subscriptionId);
  const result = await client.customers.get(billingAccountName, customerName, options);
  console.log(result);
}

customerWithExpand().catch(console.error);
