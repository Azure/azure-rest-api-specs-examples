const { AzureStackHCIClient } = require("@azure/arm-azurestackhci");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to check for updates
 *
 * @summary check for updates
 * x-ms-original-file: 2026-04-30/UpdateSummaries_CheckUpdates.json
 */
async function checkForUpdates() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "b8d594e5-51f3-4c11-9c54-a7771b81c712";
  const client = new AzureStackHCIClient(credential, subscriptionId);
  await client.updateSummariesOperationGroup.checkUpdates("testrg", "testcluster", {});
}
