const { MachineLearningWorkspacesManagementClient } = require("@azure/arm-workspaces");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Resync storage keys associated with this workspace.
 *
 * @summary Resync storage keys associated with this workspace.
 * x-ms-original-file: specification/machinelearning/resource-manager/Microsoft.MachineLearning/stable/2019-10-01/examples/ResyncStorageKeys.json
 */
async function resyncStorageKeys() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const workspaceName = "testworkspace";
  const resourceGroupName = "myResourceGroup";
  const credential = new DefaultAzureCredential();
  const client = new MachineLearningWorkspacesManagementClient(credential, subscriptionId);
  const result = await client.workspaces.resyncStorageKeys(workspaceName, resourceGroupName);
  console.log(result);
}

resyncStorageKeys().catch(console.error);
