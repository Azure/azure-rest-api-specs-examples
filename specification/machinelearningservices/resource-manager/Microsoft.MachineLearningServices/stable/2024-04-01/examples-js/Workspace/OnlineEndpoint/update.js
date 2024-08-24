const { AzureMachineLearningServicesManagementClient } = require("@azure/arm-machinelearning");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update Online Endpoint (asynchronous).
 *
 * @summary Update Online Endpoint (asynchronous).
 * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/examples/Workspace/OnlineEndpoint/update.json
 */
async function updateWorkspaceOnlineEndpoint() {
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
  const client = new AzureMachineLearningServicesManagementClient(credential, subscriptionId);
  const result = await client.onlineEndpoints.beginUpdateAndWait(
    resourceGroupName,
    workspaceName,
    endpointName,
    body,
  );
  console.log(result);
}
