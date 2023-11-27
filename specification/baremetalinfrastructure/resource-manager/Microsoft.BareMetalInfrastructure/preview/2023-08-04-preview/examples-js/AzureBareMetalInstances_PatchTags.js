const { BareMetalInfrastructureClient } = require("@azure/arm-baremetalinfrastructure");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Patches the Tags field of a Azure Bare Metal Instance for the specified subscription, resource group, and instance name.
 *
 * @summary Patches the Tags field of a Azure Bare Metal Instance for the specified subscription, resource group, and instance name.
 * x-ms-original-file: specification/baremetalinfrastructure/resource-manager/Microsoft.BareMetalInfrastructure/preview/2023-08-04-preview/examples/AzureBareMetalInstances_PatchTags.json
 */
async function updateTagsFieldOfAnAzureBareMetalInstance() {
  const subscriptionId =
    process.env["BAREMETALINFRASTRUCTURE_SUBSCRIPTION_ID"] ||
    "f0f4887f-d13c-4943-a8ba-d7da28d2a3fd";
  const resourceGroupName =
    process.env["BAREMETALINFRASTRUCTURE_RESOURCE_GROUP"] || "myResourceGroup";
  const azureBareMetalInstanceName = "myABMInstance";
  const tagsParameter = { tags: { testkey: "testvalue" } };
  const credential = new DefaultAzureCredential();
  const client = new BareMetalInfrastructureClient(credential, subscriptionId);
  const result = await client.azureBareMetalInstances.update(
    resourceGroupName,
    azureBareMetalInstanceName,
    tagsParameter
  );
  console.log(result);
}
