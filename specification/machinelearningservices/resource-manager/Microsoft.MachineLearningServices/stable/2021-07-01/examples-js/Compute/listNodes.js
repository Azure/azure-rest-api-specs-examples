const { AzureMachineLearningWorkspaces } = require("@azure/arm-machinelearning");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the details (e.g IP address, port etc) of all the compute nodes in the compute.
 *
 * @summary Get the details (e.g IP address, port etc) of all the compute nodes in the compute.
 * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2021-07-01/examples/Compute/listNodes.json
 */
async function getComputeNodesInformationForACompute() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "testrg123";
  const workspaceName = "workspaces123";
  const computeName = "compute123";
  const credential = new DefaultAzureCredential();
  const client = new AzureMachineLearningWorkspaces(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.computeOperations.listNodes(
    resourceGroupName,
    workspaceName,
    computeName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

getComputeNodesInformationForACompute().catch(console.error);
