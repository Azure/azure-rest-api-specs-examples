const { SubscriptionClient } = require("@azure/arm-resources-subscriptions");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all subscriptions for a tenant.
 *
 * @summary Gets all subscriptions for a tenant.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2022-12-01/examples/GetSubscriptions.json
 */
async function getAllSubscriptions() {
  const credential = new DefaultAzureCredential();
  const client = new SubscriptionClient(credential);
  const resArray = new Array();
  for await (let item of client.subscriptions.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}
