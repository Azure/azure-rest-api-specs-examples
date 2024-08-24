const { AzureMachineLearningServicesManagementClient } = require("@azure/arm-machinelearning");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update container.
 *
 * @summary Create or update container.
 * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/examples/Registry/DataContainer/createOrUpdate.json
 */
async function createOrUpdateRegistryDataContainer() {
  const subscriptionId =
    process.env["MACHINELEARNING_SUBSCRIPTION_ID"] || "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = process.env["MACHINELEARNING_RESOURCE_GROUP"] || "test-rg";
  const registryName = "registryName";
  const name = "string";
  const body = {
    properties: {
      description: "string",
      dataType: "uri_folder",
      isArchived: false,
      properties: { string: "string" },
      tags: { string: "string" },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureMachineLearningServicesManagementClient(credential, subscriptionId);
  const result = await client.registryDataContainers.beginCreateOrUpdateAndWait(
    resourceGroupName,
    registryName,
    name,
    body,
  );
  console.log(result);
}
