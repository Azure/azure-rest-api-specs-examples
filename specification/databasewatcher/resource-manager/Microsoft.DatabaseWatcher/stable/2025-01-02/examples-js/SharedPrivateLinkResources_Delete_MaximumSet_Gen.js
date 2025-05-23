const { DatabaseWatcherClient } = require("@azure/arm-databasewatcher");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to delete a SharedPrivateLinkResource
 *
 * @summary delete a SharedPrivateLinkResource
 * x-ms-original-file: 2025-01-02/SharedPrivateLinkResources_Delete_MaximumSet_Gen.json
 */
async function sharedPrivateLinkResourcesDeleteMaximumSet() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "49e0fbd3-75e8-44e7-96fd-5b64d9ad818d";
  const client = new DatabaseWatcherClient(credential, subscriptionId);
  await client.sharedPrivateLinkResources.delete(
    "apiTest-ddat4p",
    "databasemo3ej9ih",
    "monitoringh22eed",
  );
}
