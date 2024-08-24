const { AzureMachineLearningServicesManagementClient } = require("@azure/arm-machinelearning");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Regenerate EndpointAuthKeys for an Endpoint using Key-based authentication (asynchronous).
 *
 * @summary Regenerate EndpointAuthKeys for an Endpoint using Key-based authentication (asynchronous).
 * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/examples/Workspace/ServerlessEndpoint/regenerateKeys.json
 */
async function regenerateKeysWorkspaceServerlessEndpoint() {
  const subscriptionId =
    process.env["MACHINELEARNING_SUBSCRIPTION_ID"] || "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = process.env["MACHINELEARNING_RESOURCE_GROUP"] || "test-rg";
  const workspaceName = "my-aml-workspace";
  const name = "string";
  const body = {
    keyType: "Primary",
    keyValue: "string",
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureMachineLearningServicesManagementClient(credential, subscriptionId);
  const result = await client.serverlessEndpoints.beginRegenerateKeysAndWait(
    resourceGroupName,
    workspaceName,
    name,
    body,
  );
  console.log(result);
}
