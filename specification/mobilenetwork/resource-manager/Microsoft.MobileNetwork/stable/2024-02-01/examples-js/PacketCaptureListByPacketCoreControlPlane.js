const { MobileNetworkManagementClient } = require("@azure/arm-mobilenetwork");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all the packet capture sessions under a packet core control plane.
 *
 * @summary Lists all the packet capture sessions under a packet core control plane.
 * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2024-02-01/examples/PacketCaptureListByPacketCoreControlPlane.json
 */
async function listPacketCaptureSessionsUnderAPacketCoreControlPlane() {
  const subscriptionId =
    process.env["MOBILENETWORK_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["MOBILENETWORK_RESOURCE_GROUP"] || "rg1";
  const packetCoreControlPlaneName = "TestPacketCoreCP";
  const credential = new DefaultAzureCredential();
  const client = new MobileNetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.packetCaptures.listByPacketCoreControlPlane(
    resourceGroupName,
    packetCoreControlPlaneName,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
