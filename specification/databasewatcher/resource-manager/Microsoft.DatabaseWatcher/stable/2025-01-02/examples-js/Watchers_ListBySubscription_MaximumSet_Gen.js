const { DatabaseWatcherClient } = require("@azure/arm-databasewatcher");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to list Watcher resources by subscription ID
 *
 * @summary list Watcher resources by subscription ID
 * x-ms-original-file: 2025-01-02/Watchers_ListBySubscription_MaximumSet_Gen.json
 */
async function watchersListBySubscriptionMaximumSet() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "A76F9850-996B-40B3-94D4-C98110A0EEC9";
  const client = new DatabaseWatcherClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.watchers.listBySubscription()) {
    resArray.push(item);
  }

  console.log(resArray);
}
