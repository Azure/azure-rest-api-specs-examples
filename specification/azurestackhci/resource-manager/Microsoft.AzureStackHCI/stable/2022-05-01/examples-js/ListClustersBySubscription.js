const { AzureStackHCIClient } = require("@azure/arm-azurestackhci");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List all HCI clusters in a subscription.
 *
 * @summary List all HCI clusters in a subscription.
 * x-ms-original-file: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/stable/2022-05-01/examples/ListClustersBySubscription.json
 */
async function listClustersInAGivenSubscription() {
  const subscriptionId = "fd3c3665-1729-4b7b-9a38-238e83b0f98b";
  const credential = new DefaultAzureCredential();
  const client = new AzureStackHCIClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.clusters.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}

listClustersInAGivenSubscription().catch(console.error);
