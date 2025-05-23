const { DataBoxManagementClient } = require("@azure/arm-databox");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to This method provides the list of available skus for the given subscription, resource group and location.
 *
 * @summary This method provides the list of available skus for the given subscription, resource group and location.
 * x-ms-original-file: specification/databox/resource-manager/Microsoft.DataBox/stable/2025-02-01/examples/AvailableSkusPost.json
 */
async function availableSkusPost() {
  const subscriptionId = process.env["DATABOX_SUBSCRIPTION_ID"] || "YourSubscriptionId";
  const resourceGroupName = process.env["DATABOX_RESOURCE_GROUP"] || "YourResourceGroupName";
  const location = "westus";
  const availableSkuRequest = {
    country: "XX",
    location: "westus",
    transferType: "ImportToAzure",
  };
  const credential = new DefaultAzureCredential();
  const client = new DataBoxManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.service.listAvailableSkusByResourceGroup(
    resourceGroupName,
    location,
    availableSkuRequest,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
