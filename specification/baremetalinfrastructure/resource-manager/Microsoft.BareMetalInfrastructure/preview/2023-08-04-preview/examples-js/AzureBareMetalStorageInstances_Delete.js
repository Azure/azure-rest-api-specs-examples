const { BareMetalInfrastructureClient } = require("@azure/arm-baremetalinfrastructure");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete an AzureBareMetalStorageInstance.
 *
 * @summary Delete an AzureBareMetalStorageInstance.
 * x-ms-original-file: specification/baremetalinfrastructure/resource-manager/Microsoft.BareMetalInfrastructure/preview/2023-08-04-preview/examples/AzureBareMetalStorageInstances_Delete.json
 */
async function deleteAnAzureBareMetalStorageInstance() {
  const subscriptionId =
    process.env["BAREMETALINFRASTRUCTURE_SUBSCRIPTION_ID"] ||
    "00000000-0000-0000-0000-000000000000";
  const resourceGroupName =
    process.env["BAREMETALINFRASTRUCTURE_RESOURCE_GROUP"] || "myResourceGroup";
  const azureBareMetalStorageInstanceName = "myAzureBareMetalStorageInstance";
  const credential = new DefaultAzureCredential();
  const client = new BareMetalInfrastructureClient(credential, subscriptionId);
  const result = await client.azureBareMetalStorageInstances.delete(
    resourceGroupName,
    azureBareMetalStorageInstanceName
  );
  console.log(result);
}
