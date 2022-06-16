const { OperationalInsightsManagementClient } = require("@azure/arm-operationalinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the schema for a given workspace.
 *
 * @summary Gets the schema for a given workspace.
 * x-ms-original-file: specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2020-08-01/examples/SavedSearchesGetSchema.json
 */
async function workspacesGetSchema() {
  const subscriptionId = "00000000-0000-0000-0000-00000000000";
  const resourceGroupName = "mms-eus";
  const workspaceName = "atlantisdemo";
  const credential = new DefaultAzureCredential();
  const client = new OperationalInsightsManagementClient(credential, subscriptionId);
  const result = await client.schema.get(resourceGroupName, workspaceName);
  console.log(result);
}

workspacesGetSchema().catch(console.error);
