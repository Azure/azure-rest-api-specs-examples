const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates the properties of a billing subscription. Currently, cost center can be updated. The operation is supported only for billing accounts with agreement type Microsoft Customer Agreement.
 *
 * @summary Updates the properties of a billing subscription. Currently, cost center can be updated. The operation is supported only for billing accounts with agreement type Microsoft Customer Agreement.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/UpdateBillingSubscription.json
 */
async function updateBillingProperty() {
  const subscriptionId = "{subscriptionId}";
  const billingAccountName = "{billingAccountName}";
  const parameters = { costCenter: "ABC1234" };
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential, subscriptionId);
  const result = await client.billingSubscriptions.update(billingAccountName, parameters);
  console.log(result);
}

updateBillingProperty().catch(console.error);
