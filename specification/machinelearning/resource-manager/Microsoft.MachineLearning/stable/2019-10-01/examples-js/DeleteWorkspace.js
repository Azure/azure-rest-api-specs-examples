const { MachineLearningWorkspacesManagementClient } = require("@azure/arm-workspaces");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a machine learning workspace.
 *
 * @summary Deletes a machine learning workspace.
 * x-ms-original-file: specification/machinelearning/resource-manager/Microsoft.MachineLearning/stable/2019-10-01/examples/DeleteWorkspace.json
 */
async function workspaceDelete() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const workspaceName = "testworkspace";
  const credential = new DefaultAzureCredential();
  const client = new MachineLearningWorkspacesManagementClient(credential, subscriptionId);
  const result = await client.workspaces.delete(resourceGroupName, workspaceName);
  console.log(result);
}

workspaceDelete().catch(console.error);
