const { AzureMachineLearningWorkspaces } = require("@azure/arm-machinelearning");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a batch inference endpoint (asynchronous).
 *
 * @summary Creates a batch inference endpoint (asynchronous).
 * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-05-01/examples/BatchEndpoint/createOrUpdate.json
 */
async function createOrUpdateBatchEndpoint() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "test-rg";
  const workspaceName = "my-aml-workspace";
  const endpointName = "testEndpointName";
  const body = {
    identity: {
      type: "SystemAssigned",
      userAssignedIdentities: { string: {} },
    },
    kind: "string",
    location: "string",
    properties: {
      description: "string",
      authMode: "AMLToken",
      defaults: { deploymentName: "string" },
      properties: { string: "string" },
    },
    sku: {
      name: "string",
      capacity: 1,
      family: "string",
      size: "string",
      tier: "Free",
    },
    tags: {},
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureMachineLearningWorkspaces(credential, subscriptionId);
  const result = await client.batchEndpoints.beginCreateOrUpdateAndWait(
    resourceGroupName,
    workspaceName,
    endpointName,
    body
  );
  console.log(result);
}

createOrUpdateBatchEndpoint().catch(console.error);
