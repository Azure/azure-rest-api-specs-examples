const { MachineLearningWorkspacesManagementClient } = require("@azure/arm-workspaces");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List the authorization keys associated with this workspace.
 *
 * @summary List the authorization keys associated with this workspace.
 * x-ms-original-file: specification/machinelearning/resource-manager/Microsoft.MachineLearning/stable/2019-10-01/examples/ListWorkspaceKeys.json
 */
async function listWorkspaceKeys() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const workspaceName = "testworkspace";
  const resourceGroupName = "myResourceGroup";
  const credential = new DefaultAzureCredential();
  const client = new MachineLearningWorkspacesManagementClient(credential, subscriptionId);
  const result = await client.workspaces.listWorkspaceKeys(workspaceName, resourceGroupName);
  console.log(result);
}

listWorkspaceKeys().catch(console.error);
