const { Scvmm } = require("@azure/arm-scvmm");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List of VirtualNetworks in a subscription.
 *
 * @summary List of VirtualNetworks in a subscription.
 * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/preview/2020-06-05-preview/examples/ListVirtualNetworksBySubscription.json
 */
async function listVirtualNetworksBySubscription() {
  const subscriptionId =
    process.env["SCVMM_SUBSCRIPTION_ID"] || "fd3c3665-1729-4b7b-9a38-238e83b0f98b";
  const credential = new DefaultAzureCredential();
  const client = new Scvmm(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.virtualNetworks.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}
