const { AzureMachineLearningWorkspaces } = require("@azure/arm-machinelearning");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update a batch inference endpoint (asynchronous).
 *
 * @summary Update a batch inference endpoint (asynchronous).
 * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/BatchEndpoint/update.json
 */
async function updateBatchEndpoint() {
  const subscriptionId =
    process.env["MACHINELEARNING_SUBSCRIPTION_ID"] || "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = process.env["MACHINELEARNING_RESOURCE_GROUP"] || "test-rg";
  const workspaceName = "my-aml-workspace";
  const endpointName = "testEndpointName";
  const body = {
    identity: {
      type: "SystemAssigned",
      userAssignedIdentities: { string: {} },
    },
    tags: {},
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureMachineLearningWorkspaces(credential, subscriptionId);
  const result = await client.batchEndpoints.beginUpdateAndWait(
    resourceGroupName,
    workspaceName,
    endpointName,
    body,
  );
  console.log(result);
}
