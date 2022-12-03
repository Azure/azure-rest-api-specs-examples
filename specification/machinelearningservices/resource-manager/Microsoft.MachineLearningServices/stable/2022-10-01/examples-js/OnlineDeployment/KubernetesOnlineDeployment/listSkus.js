const { AzureMachineLearningWorkspaces } = require("@azure/arm-machinelearning");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List Inference Endpoint Deployment Skus.
 *
 * @summary List Inference Endpoint Deployment Skus.
 * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/OnlineDeployment/KubernetesOnlineDeployment/listSkus.json
 */
async function listKubernetesOnlineDeploymentSkus() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "test-rg";
  const workspaceName = "my-aml-workspace";
  const endpointName = "testEndpointName";
  const deploymentName = "testDeploymentName";
  const count = 1;
  const options = { count };
  const credential = new DefaultAzureCredential();
  const client = new AzureMachineLearningWorkspaces(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.onlineDeployments.listSkus(
    resourceGroupName,
    workspaceName,
    endpointName,
    deploymentName,
    options
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

listKubernetesOnlineDeploymentSkus().catch(console.error);
