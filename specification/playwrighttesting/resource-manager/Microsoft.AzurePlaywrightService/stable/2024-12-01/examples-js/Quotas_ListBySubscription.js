const { AzurePlaywrightServiceClient } = require("@azure/arm-playwrighttesting");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to list quotas for a given subscription Id.
 *
 * @summary list quotas for a given subscription Id.
 * x-ms-original-file: 2024-12-01/Quotas_ListBySubscription.json
 */
async function quotasListBySubscription() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new AzurePlaywrightServiceClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.quotas.listBySubscription("eastus")) {
    resArray.push(item);
  }

  console.log(resArray);
}
