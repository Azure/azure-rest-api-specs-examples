const { SubscriptionClient } = require("@azure/arm-resources-subscriptions");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets details about a specified subscription.
 *
 * @summary Gets details about a specified subscription.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2022-12-01/examples/GetSubscription.json
 */
async function getASingleSubscription() {
  const subscriptionId = "291bba3f-e0a5-47bc-a099-3bdcb2a50a05";
  const credential = new DefaultAzureCredential();
  const client = new SubscriptionClient(credential);
  const result = await client.subscriptions.get(subscriptionId);
  console.log(result);
}
