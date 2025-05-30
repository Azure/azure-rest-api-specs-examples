const { DatabaseWatcherClient } = require("@azure/arm-databasewatcher");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to list SharedPrivateLinkResource resources by Watcher
 *
 * @summary list SharedPrivateLinkResource resources by Watcher
 * x-ms-original-file: 2025-01-02/SharedPrivateLinkResources_ListByWatcher_MaximumSet_Gen.json
 */
async function sharedPrivateLinkResourcesListByWatcherMaximumSet() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "49e0fbd3-75e8-44e7-96fd-5b64d9ad818d";
  const client = new DatabaseWatcherClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.sharedPrivateLinkResources.listByWatcher(
    "apiTest-ddat4p",
    "databasemo3ej9ih",
  )) {
    resArray.push(item);
  }

  console.log(resArray);
}
