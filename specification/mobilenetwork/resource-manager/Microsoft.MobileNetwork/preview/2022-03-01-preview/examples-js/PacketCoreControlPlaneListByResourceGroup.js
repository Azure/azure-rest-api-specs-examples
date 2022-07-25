const { MobileNetworkManagementClient } = require("@azure/arm-mobilenetwork");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all the packetCoreControlPlanes in a resource group.
 *
 * @summary Lists all the packetCoreControlPlanes in a resource group.
 * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-03-01-preview/examples/PacketCoreControlPlaneListByResourceGroup.json
 */
async function listPacketCoreControlPlanesInResourceGroup() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const credential = new DefaultAzureCredential();
  const client = new MobileNetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.packetCoreControlPlanes.listByResourceGroup(resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listPacketCoreControlPlanesInResourceGroup().catch(console.error);
