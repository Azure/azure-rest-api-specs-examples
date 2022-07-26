const { SubscriptionClient } = require("@azure/arm-subscriptions");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the tenants for your account.
 *
 * @summary Gets the tenants for your account.
 * x-ms-original-file: specification/subscription/resource-manager/Microsoft.Subscription/stable/2016-06-01/examples/listTenants.json
 */
async function listTenants() {
  const credential = new DefaultAzureCredential();
  const client = new SubscriptionClient(credential);
  const resArray = new Array();
  for await (let item of client.tenants.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

listTenants().catch(console.error);
