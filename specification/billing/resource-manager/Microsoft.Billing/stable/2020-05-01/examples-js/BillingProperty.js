const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the billing properties for a subscription. This operation is not supported for billing accounts with agreement type Enterprise Agreement.
 *
 * @summary Get the billing properties for a subscription. This operation is not supported for billing accounts with agreement type Enterprise Agreement.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/BillingProperty.json
 */
async function billingProperty() {
  const subscriptionId = "{subscriptionId}";
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential, subscriptionId);
  const result = await client.billingPropertyOperations.get();
  console.log(result);
}

billingProperty().catch(console.error);
