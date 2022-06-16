const { OperationalInsightsManagementClient } = require("@azure/arm-operationalinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the specified saved search in a given workspace.
 *
 * @summary Deletes the specified saved search in a given workspace.
 * x-ms-original-file: specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2020-08-01/examples/WorkspacesDeleteSavedSearches.json
 */
async function savedSearchesDelete() {
  const subscriptionId = "00000000-0000-0000-0000-00000000000";
  const resourceGroupName = "TestRG";
  const workspaceName = "TestWS";
  const savedSearchId = "00000000-0000-0000-0000-00000000000";
  const credential = new DefaultAzureCredential();
  const client = new OperationalInsightsManagementClient(credential, subscriptionId);
  const result = await client.savedSearches.delete(resourceGroupName, workspaceName, savedSearchId);
  console.log(result);
}

savedSearchesDelete().catch(console.error);
