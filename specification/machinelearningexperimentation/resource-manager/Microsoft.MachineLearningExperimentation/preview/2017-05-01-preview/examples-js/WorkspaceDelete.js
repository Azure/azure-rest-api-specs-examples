const { MLTeamAccountManagementClient } = require("@azure/arm-machinelearningexperimentation");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a machine learning workspace.
 *
 * @summary Deletes a machine learning workspace.
 * x-ms-original-file: specification/machinelearningexperimentation/resource-manager/Microsoft.MachineLearningExperimentation/preview/2017-05-01-preview/examples/WorkspaceDelete.json
 */
async function workspaceDelete() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const accountName = "myAccount";
  const workspaceName = "testworkspace";
  const credential = new DefaultAzureCredential();
  const client = new MLTeamAccountManagementClient(credential, subscriptionId);
  const result = await client.workspaces.delete(resourceGroupName, accountName, workspaceName);
  console.log(result);
}

workspaceDelete().catch(console.error);
