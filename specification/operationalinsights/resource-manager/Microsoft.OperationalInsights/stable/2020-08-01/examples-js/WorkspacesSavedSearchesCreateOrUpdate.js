const { OperationalInsightsManagementClient } = require("@azure/arm-operationalinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a saved search for a given workspace.
 *
 * @summary Creates or updates a saved search for a given workspace.
 * x-ms-original-file: specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2020-08-01/examples/WorkspacesSavedSearchesCreateOrUpdate.json
 */
async function savedSearchCreateOrUpdate() {
  const subscriptionId =
    process.env["OPERATIONALINSIGHTS_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-00000000000";
  const resourceGroupName = process.env["OPERATIONALINSIGHTS_RESOURCE_GROUP"] || "TestRG";
  const workspaceName = "TestWS";
  const savedSearchId = "00000000-0000-0000-0000-00000000000";
  const parameters = {
    category: "Saved Search Test Category",
    displayName: "Create or Update Saved Search Test",
    functionAlias: "heartbeat_func",
    functionParameters: "a:int=1",
    query: "Heartbeat | summarize Count() by Computer | take a",
    tags: [{ name: "Group", value: "Computer" }],
    version: 2,
  };
  const credential = new DefaultAzureCredential();
  const client = new OperationalInsightsManagementClient(credential, subscriptionId);
  const result = await client.savedSearches.createOrUpdate(
    resourceGroupName,
    workspaceName,
    savedSearchId,
    parameters
  );
  console.log(result);
}
