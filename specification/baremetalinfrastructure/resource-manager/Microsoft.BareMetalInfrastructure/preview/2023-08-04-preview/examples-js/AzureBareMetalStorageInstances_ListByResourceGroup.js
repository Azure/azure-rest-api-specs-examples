const { BareMetalInfrastructureClient } = require("@azure/arm-baremetalinfrastructure");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a list of AzureBareMetalStorage instances in the specified subscription and resource group. The operations returns various properties of each Azure Bare Metal Instance.
 *
 * @summary Gets a list of AzureBareMetalStorage instances in the specified subscription and resource group. The operations returns various properties of each Azure Bare Metal Instance.
 * x-ms-original-file: specification/baremetalinfrastructure/resource-manager/Microsoft.BareMetalInfrastructure/preview/2023-08-04-preview/examples/AzureBareMetalStorageInstances_ListByResourceGroup.json
 */
async function listAllAzureBareMetalStorageInstancesInAResourceGroup() {
  const subscriptionId =
    process.env["BAREMETALINFRASTRUCTURE_SUBSCRIPTION_ID"] ||
    "f0f4887f-d13c-4943-a8ba-d7da28d2a3fd";
  const resourceGroupName =
    process.env["BAREMETALINFRASTRUCTURE_RESOURCE_GROUP"] || "myResourceGroup";
  const credential = new DefaultAzureCredential();
  const client = new BareMetalInfrastructureClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.azureBareMetalStorageInstances.listByResourceGroup(
    resourceGroupName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
