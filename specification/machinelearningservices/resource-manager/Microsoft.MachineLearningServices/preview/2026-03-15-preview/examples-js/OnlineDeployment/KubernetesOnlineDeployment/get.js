const { AzureMachineLearningServicesManagementClient } = require("@azure/arm-machinelearning");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to get Inference Deployment Deployment.
 *
 * @summary get Inference Deployment Deployment.
 * x-ms-original-file: 2026-03-15-preview/OnlineDeployment/KubernetesOnlineDeployment/get.json
 */
async function getKubernetesOnlineDeployment() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const client = new AzureMachineLearningServicesManagementClient(credential, subscriptionId);
  const result = await client.onlineDeployments.get(
    "test-rg",
    "my-aml-workspace",
    "testEndpointName",
    "testDeploymentName",
  );
  console.log(result);
}
