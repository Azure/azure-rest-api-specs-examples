const { SubscriptionClient } = require("@azure/arm-subscriptions");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get Alias Subscription.
 *
 * @summary Get Alias Subscription.
 * x-ms-original-file: specification/subscription/resource-manager/Microsoft.Subscription/stable/2021-10-01/examples/getAlias.json
 */
async function getAlias() {
  const aliasName = "aliasForNewSub";
  const credential = new DefaultAzureCredential();
  const client = new SubscriptionClient(credential);
  const result = await client.alias.get(aliasName);
  console.log(result);
}

getAlias().catch(console.error);
