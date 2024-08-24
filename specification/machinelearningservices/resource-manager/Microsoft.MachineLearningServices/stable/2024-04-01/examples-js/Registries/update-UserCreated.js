const { AzureMachineLearningServicesManagementClient } = require("@azure/arm-machinelearning");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update tags
 *
 * @summary Update tags
 * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/examples/Registries/update-UserCreated.json
 */
async function updateRegistryWithUserCreatedAccounts() {
  const subscriptionId =
    process.env["MACHINELEARNING_SUBSCRIPTION_ID"] || "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = process.env["MACHINELEARNING_RESOURCE_GROUP"] || "test-rg";
  const registryName = "string";
  const body = {
    identity: { type: "UserAssigned", userAssignedIdentities: { string: {} } },
    sku: {
      name: "string",
      capacity: 1,
      family: "string",
      size: "string",
      tier: "Basic",
    },
    tags: {},
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureMachineLearningServicesManagementClient(credential, subscriptionId);
  const result = await client.registries.update(resourceGroupName, registryName, body);
  console.log(result);
}
