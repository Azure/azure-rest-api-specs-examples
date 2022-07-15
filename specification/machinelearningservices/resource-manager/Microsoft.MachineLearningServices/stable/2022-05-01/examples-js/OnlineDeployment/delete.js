const { AzureMachineLearningWorkspaces } = require("@azure/arm-machinelearning");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete Inference Endpoint Deployment (asynchronous).
 *
 * @summary Delete Inference Endpoint Deployment (asynchronous).
 * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-05-01/examples/OnlineDeployment/delete.json
 */
async function deleteOnlineDeployment() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "testrg123";
  const workspaceName = "workspace123";
  const endpointName = "testEndpoint";
  const deploymentName = "testDeployment";
  const credential = new DefaultAzureCredential();
  const client = new AzureMachineLearningWorkspaces(credential, subscriptionId);
  const result = await client.onlineDeployments.beginDeleteAndWait(
    resourceGroupName,
    workspaceName,
    endpointName,
    deploymentName
  );
  console.log(result);
}

deleteOnlineDeployment().catch(console.error);
