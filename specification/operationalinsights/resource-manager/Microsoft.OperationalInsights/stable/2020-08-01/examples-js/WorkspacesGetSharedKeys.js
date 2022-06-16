const { OperationalInsightsManagementClient } = require("@azure/arm-operationalinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the shared keys for a workspace.
 *
 * @summary Gets the shared keys for a workspace.
 * x-ms-original-file: specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2020-08-01/examples/WorkspacesGetSharedKeys.json
 */
async function sharedKeysList() {
  const subscriptionId = "00000000-0000-0000-0000-00000000000";
  const resourceGroupName = "rg1";
  const workspaceName = "TestLinkWS";
  const credential = new DefaultAzureCredential();
  const client = new OperationalInsightsManagementClient(credential, subscriptionId);
  const result = await client.sharedKeysOperations.getSharedKeys(resourceGroupName, workspaceName);
  console.log(result);
}

sharedKeysList().catch(console.error);
