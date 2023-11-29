const { MobileNetworkManagementClient } = require("@azure/arm-mobilenetwork");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all the packet core data planes associated with a packet core control plane.
 *
 * @summary Lists all the packet core data planes associated with a packet core control plane.
 * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2023-09-01/examples/PacketCoreDataPlaneListByPacketCoreControlPlane.json
 */
async function listPacketCoreDataPlanesInAControlPlane() {
  const subscriptionId =
    process.env["MOBILENETWORK_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["MOBILENETWORK_RESOURCE_GROUP"] || "rg1";
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
