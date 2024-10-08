const { AzureMachineLearningServicesManagementClient } = require("@azure/arm-machinelearning");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List all the secrets of a machine learning workspaces connections.
 *
 * @summary List all the secrets of a machine learning workspaces connections.
 * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/examples/WorkspaceConnection/listSecrets.json
 */
async function getWorkspaceConnection() {
  const subscriptionId =
    process.env["MACHINELEARNING_SUBSCRIPTION_ID"] || "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = process.env["MACHINELEARNING_RESOURCE_GROUP"] || "test-rg";
  const workspaceName = "workspace-1";
  const connectionName = "connection-1";
  const credential = new DefaultAzureCredential();
  const client = new AzureMachineLearningServicesManagementClient(credential, subscriptionId);
  const result = await client.workspaceConnections.listSecrets(
    resourceGroupName,
    workspaceName,
    connectionName,
  );
  console.log(result);
}
