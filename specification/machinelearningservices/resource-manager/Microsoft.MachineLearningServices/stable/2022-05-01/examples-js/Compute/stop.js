const { AzureMachineLearningWorkspaces } = require("@azure/arm-machinelearning");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Posts a stop action to a compute instance
 *
 * @summary Posts a stop action to a compute instance
 * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-05-01/examples/Compute/stop.json
 */
async function stopComputeInstanceCompute() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "testrg123";
  const workspaceName = "workspaces123";
  const computeName = "compute123";
  const credential = new DefaultAzureCredential();
  const client = new AzureMachineLearningWorkspaces(credential, subscriptionId);
  const result = await client.computeOperations.beginStopAndWait(
    resourceGroupName,
    workspaceName,
    computeName
  );
  console.log(result);
}

stopComputeInstanceCompute().catch(console.error);
