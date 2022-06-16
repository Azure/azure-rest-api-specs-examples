const { AzureMachineLearningWorkspaces } = require("@azure/arm-machinelearning");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Add a new workspace connection.
 *
 * @summary Add a new workspace connection.
 * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2021-07-01/examples/WorkspaceConnection/create.json
 */
async function createWorkspaceConnection() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "resourceGroup-1";
  const workspaceName = "workspace-1";
  const connectionName = "connection-1";
  const parameters = {
    authType: "PAT",
    category: "ACR",
    target: "www.facebook.com",
    value: "secrets",
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureMachineLearningWorkspaces(credential, subscriptionId);
  const result = await client.workspaceConnections.create(
    resourceGroupName,
    workspaceName,
    connectionName,
    parameters
  );
  console.log(result);
}

createWorkspaceConnection().catch(console.error);
