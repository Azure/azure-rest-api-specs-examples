const { StorageClient } = require("@azure/arm-dell-storage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to list FileSystemResource resources by subscription ID
 *
 * @summary list FileSystemResource resources by subscription ID
 * x-ms-original-file: 2025-03-21/FileSystems_ListBySubscription_MaximumSet_Gen.json
 */
async function fileSystemsListBySubscriptionMaximumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "4B6E265D-57CF-4A9D-8B35-3CC68ED9D208";
  const client = new StorageClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.fileSystems.listBySubscription()) {
    resArray.push(item);
  }

  console.log(resArray);
}
