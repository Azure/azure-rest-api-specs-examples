const { AzureMachineLearningServicesManagementClient } = require("@azure/arm-machinelearning");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete Batch Inference Endpoint (asynchronous).
 *
 * @summary Delete Batch Inference Endpoint (asynchronous).
 * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/examples/Workspace/BatchEndpoint/delete.json
 */
async function deleteWorkspaceBatchEndpoint() {
  const subscriptionId =
    process.env["MACHINELEARNING_SUBSCRIPTION_ID"] || "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = process.env["MACHINELEARNING_RESOURCE_GROUP"] || "resourceGroup-1234";
  const workspaceName = "testworkspace";
  const endpointName = "testBatchEndpoint";
  const credential = new DefaultAzureCredential();
  const client = new AzureMachineLearningServicesManagementClient(credential, subscriptionId);
  const result = await client.batchEndpoints.beginDeleteAndWait(
    resourceGroupName,
    workspaceName,
    endpointName,
  );
  console.log(result);
}
