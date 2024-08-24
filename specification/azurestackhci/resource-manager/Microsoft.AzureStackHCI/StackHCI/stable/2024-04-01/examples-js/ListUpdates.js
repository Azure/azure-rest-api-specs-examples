const { AzureStackHCIClient } = require("@azure/arm-azurestackhci");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List all Updates
 *
 * @summary List all Updates
 * x-ms-original-file: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/StackHCI/stable/2024-04-01/examples/ListUpdates.json
 */
async function listAvailableUpdates() {
  const subscriptionId =
    process.env["AZURESTACKHCI_SUBSCRIPTION_ID"] || "b8d594e5-51f3-4c11-9c54-a7771b81c712";
  const resourceGroupName = process.env["AZURESTACKHCI_RESOURCE_GROUP"] || "testrg";
  const clusterName = "testcluster";
  const credential = new DefaultAzureCredential();
  const client = new AzureStackHCIClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.updates.list(resourceGroupName, clusterName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
