const { MobileNetworkManagementClient } = require("@azure/arm-mobilenetwork");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates a PacketCoreControlPlane update tags.
 *
 * @summary Updates a PacketCoreControlPlane update tags.
 * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-03-01-preview/examples/PacketCoreControlPlaneUpdateTags.json
 */
async function updatePacketCoreControlPlaneTags() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const packetCoreControlPlaneName = "TestPacketCoreCP";
  const parameters = { tags: { tag1: "value1", tag2: "value2" } };
  const credential = new DefaultAzureCredential();
  const client = new MobileNetworkManagementClient(credential, subscriptionId);
  const result = await client.packetCoreControlPlanes.updateTags(
    resourceGroupName,
    packetCoreControlPlaneName,
    parameters
  );
  console.log(result);
}

updatePacketCoreControlPlaneTags().catch(console.error);
