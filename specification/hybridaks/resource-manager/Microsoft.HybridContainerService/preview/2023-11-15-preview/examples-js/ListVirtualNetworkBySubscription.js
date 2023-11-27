const { HybridContainerServiceClient } = require("@azure/arm-hybridcontainerservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the Hybrid AKS virtual networks by subscription
 *
 * @summary Lists the Hybrid AKS virtual networks by subscription
 * x-ms-original-file: specification/hybridaks/resource-manager/Microsoft.HybridContainerService/preview/2023-11-15-preview/examples/ListVirtualNetworkBySubscription.json
 */
async function listVirtualNetworkBySubscription() {
  const subscriptionId =
    process.env["HYBRIDCONTAINERSERVICE_SUBSCRIPTION_ID"] || "a3e42606-29b1-4d7d-b1d9-9ff6b9d3c71b";
  const credential = new DefaultAzureCredential();
  const client = new HybridContainerServiceClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.virtualNetworks.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}
