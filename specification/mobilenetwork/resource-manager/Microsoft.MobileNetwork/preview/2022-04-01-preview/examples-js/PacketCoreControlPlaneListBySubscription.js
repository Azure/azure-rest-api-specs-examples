const { MobileNetworkManagementClient } = require("@azure/arm-mobilenetwork");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all the packet core control planes in a subscription.
 *
 * @summary Lists all the packet core control planes in a subscription.
 * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-04-01-preview/examples/PacketCoreControlPlaneListBySubscription.json
 */
async function listPacketCoreControlPlanesInASubscription() {
  const subscriptionId = "subid";
  const credential = new DefaultAzureCredential();
  const client = new MobileNetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.packetCoreControlPlanes.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}

listPacketCoreControlPlanesInASubscription().catch(console.error);
