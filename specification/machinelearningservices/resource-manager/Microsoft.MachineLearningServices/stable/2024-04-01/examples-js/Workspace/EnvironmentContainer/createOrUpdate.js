const { AzureMachineLearningServicesManagementClient } = require("@azure/arm-machinelearning");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update container.
 *
 * @summary Create or update container.
 * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/examples/Workspace/EnvironmentContainer/createOrUpdate.json
 */
async function createOrUpdateWorkspaceEnvironmentContainer() {
  const subscriptionId =
    process.env["MACHINELEARNING_SUBSCRIPTION_ID"] || "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = process.env["MACHINELEARNING_RESOURCE_GROUP"] || "testrg123";
  const workspaceName = "testworkspace";
  const name = "testEnvironment";
  const body = {
    properties: {
      description: "string",
      properties: {
        additionalProp1: "string",
        additionalProp2: "string",
        additionalProp3: "string",
      },
      tags: {
        additionalProp1: "string",
        additionalProp2: "string",
        additionalProp3: "string",
      },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureMachineLearningServicesManagementClient(credential, subscriptionId);
  const result = await client.environmentContainers.createOrUpdate(
    resourceGroupName,
    workspaceName,
    name,
    body,
  );
  console.log(result);
}
