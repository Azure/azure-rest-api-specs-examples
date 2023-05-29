const { SubscriptionClient } = require("@azure/arm-resources-subscriptions");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the tenants for your account.
 *
 * @summary Gets the tenants for your account.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2022-12-01/examples/GetTenants.json
 */
async function getAllTenants() {
  const credential = new DefaultAzureCredential();
  const client = new SubscriptionClient(credential);
  const resArray = new Array();
  for await (let item of client.tenants.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}
