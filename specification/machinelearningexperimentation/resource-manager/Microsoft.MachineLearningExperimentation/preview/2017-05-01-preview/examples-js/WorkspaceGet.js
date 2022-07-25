const { MLTeamAccountManagementClient } = require("@azure/arm-machinelearningexperimentation");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the properties of the specified machine learning workspace.
 *
 * @summary Gets the properties of the specified machine learning workspace.
 * x-ms-original-file: specification/machinelearningexperimentation/resource-manager/Microsoft.MachineLearningExperimentation/preview/2017-05-01-preview/examples/WorkspaceGet.json
 */
async function workspaceGet() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "accountcrud-1234";
  const accountName = "accountcrud5678";
  const workspaceName = "testworkspace";
  const credential = new DefaultAzureCredential();
  const client = new MLTeamAccountManagementClient(credential, subscriptionId);
  const result = await client.workspaces.get(resourceGroupName, accountName, workspaceName);
  console.log(result);
}

workspaceGet().catch(console.error);
