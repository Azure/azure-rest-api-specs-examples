const { MobileNetworkManagementClient } = require("@azure/arm-mobilenetwork");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all the packet core data planes associated with a packet core control plane.
 *
 * @summary Lists all the packet core data planes associated with a packet core control plane.
 * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-04-01-preview/examples/PacketCoreDataPlaneListByPacketCoreControlPlane.json
 */
async function listPacketCoreDataPlanesInAControlPlane() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const packetCoreControlPlaneName = "testPacketCoreCP";
  const credential = new DefaultAzureCredential();
  const client = new MobileNetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.packetCoreDataPlanes.listByPacketCoreControlPlane(
    resourceGroupName,
    packetCoreControlPlaneName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

listPacketCoreDataPlanesInAControlPlane().catch(console.error);
