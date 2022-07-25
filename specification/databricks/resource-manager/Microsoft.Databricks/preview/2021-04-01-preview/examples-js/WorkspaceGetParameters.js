const { AzureDatabricksManagementClient } = require("@azure/arm-databricks");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the workspace.
 *
 * @summary Gets the workspace.
 * x-ms-original-file: specification/databricks/resource-manager/Microsoft.Databricks/preview/2021-04-01-preview/examples/WorkspaceGetParameters.json
 */
async function getAWorkspaceWithCustomParameters() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg";
  const workspaceName = "myWorkspace";
  const credential = new DefaultAzureCredential();
  const client = new AzureDatabricksManagementClient(credential, subscriptionId);
  const result = await client.workspaces.get(resourceGroupName, workspaceName);
  console.log(result);
}

getAWorkspaceWithCustomParameters().catch(console.error);
