const { MachineLearningComputeManagementClient } = require("@azure/arm-machinelearningcompute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all available operations.
 *
 * @summary Gets all available operations.
 * x-ms-original-file: specification/machinelearningcompute/resource-manager/Microsoft.MachineLearningCompute/preview/2017-08-01-preview/examples/MachineLearningCompute_ListAvailableOperations.json
 */
async function machineLearningComputeListAvailableOperations() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new MachineLearningComputeManagementClient(credential, subscriptionId);
  const result = await client.machineLearningCompute.listAvailableOperations();
  console.log(result);
}

machineLearningComputeListAvailableOperations().catch(console.error);
