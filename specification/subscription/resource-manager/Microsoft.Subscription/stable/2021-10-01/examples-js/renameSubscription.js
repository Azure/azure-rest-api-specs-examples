const { SubscriptionClient } = require("@azure/arm-subscriptions");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The operation to rename a subscription
 *
 * @summary The operation to rename a subscription
 * x-ms-original-file: specification/subscription/resource-manager/Microsoft.Subscription/stable/2021-10-01/examples/renameSubscription.json
 */
async function renameSubscription() {
  const subscriptionId = "83aa47df-e3e9-49ff-877b-94304bf3d3ad";
  const body = { subscriptionName: "Test Sub" };
  const credential = new DefaultAzureCredential();
  const client = new SubscriptionClient(credential);
  const result = await client.subscriptionOperations.rename(subscriptionId, body);
  console.log(result);
}

renameSubscription().catch(console.error);
