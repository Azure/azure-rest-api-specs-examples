const { DataBoxManagementClient } = require("@azure/arm-databox");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to This method provides the list of available skus for the given subscription, resource group and location.
 *
 * @summary This method provides the list of available skus for the given subscription, resource group and location.
 * x-ms-original-file: specification/databox/resource-manager/Microsoft.DataBox/preview/2021-08-01-preview/examples/AvailableSkusPost.json
 */
async function availableSkusPost() {
  const subscriptionId = "fa68082f-8ff7-4a25-95c7-ce9da541242f";
  const resourceGroupName = "bvttoolrg6";
  const location = "westus";
  const availableSkuRequest = {
    country: "US",
    location: "westus",
    transferType: "ImportToAzure",
  };
  const credential = new DefaultAzureCredential();
  const client = new DataBoxManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.service.listAvailableSkusByResourceGroup(
    resourceGroupName,
    location,
    availableSkuRequest
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

availableSkusPost().catch(console.error);
