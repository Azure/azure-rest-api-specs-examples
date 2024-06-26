const { PlaywrightTestingClient } = require("@azure/arm-playwrighttesting");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List quotas for a given subscription Id.
 *
 * @summary List quotas for a given subscription Id.
 * x-ms-original-file: specification/playwrighttesting/resource-manager/Microsoft.AzurePlaywrightService/preview/2023-10-01-preview/examples/Quotas_ListBySubscription.json
 */
async function quotasListBySubscription() {
  const subscriptionId =
    process.env["PLAYWRIGHTTESTING_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const location = "eastus";
  const credential = new DefaultAzureCredential();
  const client = new PlaywrightTestingClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.quotas.listBySubscription(location)) {
    resArray.push(item);
  }
  console.log(resArray);
}
