const { MobileNetworkManagementClient } = require("@azure/arm-mobilenetwork");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates a PacketCoreDataPlane update tags.
 *
 * @summary Updates a PacketCoreDataPlane update tags.
 * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-03-01-preview/examples/PacketCoreDataPlaneUpdateTags.json
 */
async function updatePacketCoreDataPlaneTags() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const packetCoreControlPlaneName = "testPacketCoreCP";
  const packetCoreDataPlaneName = "testPacketCoreDP";
  const parameters = { tags: { tag1: "value1", tag2: "value2" } };
  const credential = new DefaultAzureCredential();
  const client = new MobileNetworkManagementClient(credential, subscriptionId);
  const result = await client.packetCoreDataPlanes.updateTags(
    resourceGroupName,
    packetCoreControlPlaneName,
    packetCoreDataPlaneName,
    parameters
  );
  console.log(result);
}

updatePacketCoreDataPlaneTags().catch(console.error);
