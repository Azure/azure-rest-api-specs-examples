const { AzureMachineLearningWorkspaces } = require("@azure/arm-machinelearning");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a batch inference endpoint by name.
 *
 * @summary Gets a batch inference endpoint by name.
 * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-05-01/examples/BatchEndpoint/get.json
 */
async function getBatchEndpoint() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "test-rg";
  const workspaceName = "my-aml-workspace";
  const endpointName = "testEndpointName";
  const credential = new DefaultAzureCredential();
  const client = new AzureMachineLearningWorkspaces(credential, subscriptionId);
  const result = await client.batchEndpoints.get(resourceGroupName, workspaceName, endpointName);
  console.log(result);
}

getBatchEndpoint().catch(console.error);
