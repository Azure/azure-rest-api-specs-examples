const { SubscriptionClient } = require("@azure/arm-subscriptions");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The operation to enable a subscription
 *
 * @summary The operation to enable a subscription
 * x-ms-original-file: specification/subscription/resource-manager/Microsoft.Subscription/stable/2021-10-01/examples/enableSubscription.json
 */
async function enableSubscription() {
  const subscriptionId = "7948bcee-488c-47ce-941c-38e20ede803d";
  const credential = new DefaultAzureCredential();
  const client = new SubscriptionClient(credential);
  const result = await client.subscriptionOperations.enable(subscriptionId);
  console.log(result);
}

enableSubscription().catch(console.error);
