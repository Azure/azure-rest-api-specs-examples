const { SubscriptionClient } = require("@azure/arm-subscriptions");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the subscription tenant policy for the user's tenant.
 *
 * @summary Get the subscription tenant policy for the user's tenant.
 * x-ms-original-file: specification/subscription/resource-manager/Microsoft.Subscription/stable/2021-10-01/examples/getTenantPolicy.json
 */
async function getTenantPolicy() {
  const credential = new DefaultAzureCredential();
  const client = new SubscriptionClient(credential);
  const result = await client.subscriptionPolicy.getPolicyForTenant();
  console.log(result);
}

getTenantPolicy().catch(console.error);
