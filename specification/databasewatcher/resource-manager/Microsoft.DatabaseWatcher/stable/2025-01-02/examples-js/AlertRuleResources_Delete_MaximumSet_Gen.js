const { DatabaseWatcherClient } = require("@azure/arm-databasewatcher");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to delete a AlertRuleResource
 *
 * @summary delete a AlertRuleResource
 * x-ms-original-file: 2025-01-02/AlertRuleResources_Delete_MaximumSet_Gen.json
 */
async function alertRuleResourcesDeleteMaximumSet() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "A76F9850-996B-40B3-94D4-C98110A0EEC9";
  const client = new DatabaseWatcherClient(credential, subscriptionId);
  await client.alertRuleResources.delete("rgWatcher", "testWatcher", "testAlert");
}
