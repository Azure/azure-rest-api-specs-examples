const { BareMetalInfrastructureClient } = require("@azure/arm-baremetalinfrastructure");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets an Azure Bare Metal Instance for the specified subscription, resource group, and instance name.
 *
 * @summary Gets an Azure Bare Metal Instance for the specified subscription, resource group, and instance name.
 * x-ms-original-file: specification/baremetalinfrastructure/resource-manager/Microsoft.BareMetalInfrastructure/preview/2023-08-04-preview/examples/AzureBareMetalInstances_Get.json
 */
async function getAnAzureBareMetalInstance() {
  const subscriptionId =
    process.env["BAREMETALINFRASTRUCTURE_SUBSCRIPTION_ID"] ||
    "f0f4887f-d13c-4943-a8ba-d7da28d2a3fd";
  const resourceGroupName =
    process.env["BAREMETALINFRASTRUCTURE_RESOURCE_GROUP"] || "myResourceGroup";
  const azureBareMetalInstanceName = "myAzureBareMetalInstance";
  const credential = new DefaultAzureCredential();
  const client = new BareMetalInfrastructureClient(credential, subscriptionId);
  const result = await client.azureBareMetalInstances.get(
    resourceGroupName,
    azureBareMetalInstanceName
  );
  console.log(result);
}
