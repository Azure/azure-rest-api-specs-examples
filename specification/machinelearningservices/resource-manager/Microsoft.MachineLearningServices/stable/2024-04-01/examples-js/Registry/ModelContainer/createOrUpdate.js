const { AzureMachineLearningServicesManagementClient } = require("@azure/arm-machinelearning");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update model container.
 *
 * @summary Create or update model container.
 * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/examples/Registry/ModelContainer/createOrUpdate.json
 */
async function createOrUpdateRegistryModelContainer() {
  const subscriptionId =
    process.env["MACHINELEARNING_SUBSCRIPTION_ID"] || "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = process.env["MACHINELEARNING_RESOURCE_GROUP"] || "testrg123";
  const registryName = "registry123";
  const modelName = "testContainer";
  const body = {
    properties: {
      description: "Model container description",
      tags: { tag1: "value1", tag2: "value2" },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureMachineLearningServicesManagementClient(credential, subscriptionId);
  const result = await client.registryModelContainers.beginCreateOrUpdateAndWait(
    resourceGroupName,
    registryName,
    modelName,
    body,
  );
  console.log(result);
}
