const { MobileNetworkManagementClient } = require("@azure/arm-mobilenetwork");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all the mobile networks in a subscription.
 *
 * @summary Lists all the mobile networks in a subscription.
 * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-04-01-preview/examples/MobileNetworkListBySubscription.json
 */
async function listMobileNetworksInASubscription() {
  const subscriptionId = "subid";
  const credential = new DefaultAzureCredential();
  const client = new MobileNetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.mobileNetworks.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}

listMobileNetworksInASubscription().catch(console.error);
