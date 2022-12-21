const { MobileNetworkManagementClient } = require("@azure/arm-mobilenetwork");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates packet core control planes tags.
 *
 * @summary Updates packet core control planes tags.
 * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2022-11-01/examples/PacketCoreControlPlaneUpdateTags.json
 */
async function updatePacketCoreControlPlaneTags() {
  const subscriptionId = process.env["MOBILENETWORK_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["MOBILENETWORK_RESOURCE_GROUP"] || "rg1";
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
