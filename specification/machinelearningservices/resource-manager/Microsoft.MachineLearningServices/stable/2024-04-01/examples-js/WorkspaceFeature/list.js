const { AzureMachineLearningServicesManagementClient } = require("@azure/arm-machinelearning");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all enabled features for a workspace
 *
 * @summary Lists all enabled features for a workspace
 * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/examples/WorkspaceFeature/list.json
 */
async function listWorkspaceFeatures() {
  const subscriptionId =
    process.env["MACHINELEARNING_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["MACHINELEARNING_RESOURCE_GROUP"] || "myResourceGroup";
  const workspaceName = "testworkspace";
  const credential = new DefaultAzureCredential();
  const client = new AzureMachineLearningServicesManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.workspaceFeatures.list(resourceGroupName, workspaceName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
