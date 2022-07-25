const { MobileNetworkManagementClient } = require("@azure/arm-mobilenetwork");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the specified packet core data plane.
 *
 * @summary Deletes the specified packet core data plane.
 * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-03-01-preview/examples/PacketCoreDataPlaneDelete.json
 */
async function deletePacketCoreDataPlane() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const packetCoreControlPlaneName = "testPacketCoreCP";
  const packetCoreDataPlaneName = "testPacketCoreDP";
  const credential = new DefaultAzureCredential();
  const client = new MobileNetworkManagementClient(credential, subscriptionId);
  const result = await client.packetCoreDataPlanes.beginDeleteAndWait(
    resourceGroupName,
    packetCoreControlPlaneName,
    packetCoreDataPlaneName
  );
  console.log(result);
}

deletePacketCoreDataPlane().catch(console.error);
