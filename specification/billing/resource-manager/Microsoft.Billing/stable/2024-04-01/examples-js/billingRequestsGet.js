const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a billing request by its ID.
 *
 * @summary Gets a billing request by its ID.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/billingRequestsGet.json
 */
async function billingRequestsGet() {
  const billingRequestName = "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential);
  const result = await client.billingRequests.get(billingRequestName);
  console.log(result);
}
