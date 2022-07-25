const { SubscriptionClient } = require("@azure/arm-subscriptions");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all subscriptions for a tenant.
 *
 * @summary Gets all subscriptions for a tenant.
 * x-ms-original-file: specification/subscription/resource-manager/Microsoft.Subscription/stable/2016-06-01/examples/listSubscriptions.json
 */
async function listSubscriptions() {
  const credential = new DefaultAzureCredential();
  const client = new SubscriptionClient(credential);
  const resArray = new Array();
  for await (let item of client.subscriptions.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

listSubscriptions().catch(console.error);
