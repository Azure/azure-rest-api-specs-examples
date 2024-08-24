const { AzureMachineLearningServicesManagementClient } = require("@azure/arm-machinelearning");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all the keys associated with this workspace. This includes keys for the storage account, app insights and password for container registry
 *
 * @summary Lists all the keys associated with this workspace. This includes keys for the storage account, app insights and password for container registry
 * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/examples/Workspace/listKeys.json
 */
async function listWorkspaceKeys() {
  const subscriptionId =
    process.env["MACHINELEARNING_SUBSCRIPTION_ID"] || "aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee";
  const resourceGroupName = process.env["MACHINELEARNING_RESOURCE_GROUP"] || "testrg123";
  const workspaceName = "workspaces123";
  const credential = new DefaultAzureCredential();
  const client = new AzureMachineLearningServicesManagementClient(credential, subscriptionId);
  const result = await client.workspaces.listKeys(resourceGroupName, workspaceName);
  console.log(result);
}
