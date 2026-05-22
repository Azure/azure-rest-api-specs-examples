const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to gets the billing properties for a subscription
 *
 * @summary gets the billing properties for a subscription
 * x-ms-original-file: 2024-04-01/billingPropertyGetMCA.json
 */
async function billingPropertyGetMCA() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "11111111-1111-1111-1111-111111111111";
  const client = new BillingManagementClient(credential, subscriptionId);
  const result = await client.billingProperty.get();
  console.log(result);
}
