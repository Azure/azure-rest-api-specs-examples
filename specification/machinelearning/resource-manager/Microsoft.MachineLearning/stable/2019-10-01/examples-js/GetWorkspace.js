const { MachineLearningWorkspacesManagementClient } = require("@azure/arm-workspaces");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the properties of the specified machine learning workspace.
 *
 * @summary Gets the properties of the specified machine learning workspace.
 * x-ms-original-file: specification/machinelearning/resource-manager/Microsoft.MachineLearning/stable/2019-10-01/examples/GetWorkspace.json
 */
async function workspaceGet() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const workspaceName = "testworkspace";
  const credential = new DefaultAzureCredential();
  const client = new MachineLearningWorkspacesManagementClient(credential, subscriptionId);
  const result = await client.workspaces.get(resourceGroupName, workspaceName);
  console.log(result);
}

workspaceGet().catch(console.error);
