const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to list network usages for a subscription.
 *
 * @summary list network usages for a subscription.
 * x-ms-original-file: 2025-07-01/UsageList.json
 */
async function listUsages() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.usages.list("westus")) {
    resArray.push(item);
  }

  console.log(resArray);
}
