const { HybridContainerServiceClient } = require("@azure/arm-hybridcontainerservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the virtual networks in the specified subscription
 *
 * @summary Lists the virtual networks in the specified subscription
 * x-ms-original-file: specification/hybridaks/resource-manager/Microsoft.HybridContainerService/stable/2024-01-01/examples/ListVirtualNetworkBySubscription.json
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
