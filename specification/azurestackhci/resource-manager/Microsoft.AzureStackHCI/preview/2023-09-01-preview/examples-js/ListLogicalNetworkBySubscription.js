const { AzureStackHCIClient } = require("@azure/arm-azurestackhci");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all of the logical networks in the specified subscription. Use the nextLink property in the response to get the next page of logical networks.
 *
 * @summary Lists all of the logical networks in the specified subscription. Use the nextLink property in the response to get the next page of logical networks.
 * x-ms-original-file: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/preview/2023-09-01-preview/examples/ListLogicalNetworkBySubscription.json
 */
async function listLogicalNetworkBySubscription() {
  const subscriptionId =
    process.env["AZURESTACKHCI_SUBSCRIPTION_ID"] || "fd3c3665-1729-4b7b-9a38-238e83b0f98b";
  const credential = new DefaultAzureCredential();
  const client = new AzureStackHCIClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.logicalNetworksOperations.listAll()) {
    resArray.push(item);
  }
  console.log(resArray);
}
