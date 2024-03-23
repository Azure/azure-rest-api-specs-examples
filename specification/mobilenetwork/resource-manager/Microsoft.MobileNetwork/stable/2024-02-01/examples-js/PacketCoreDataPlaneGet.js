const { MobileNetworkManagementClient } = require("@azure/arm-mobilenetwork");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets information about the specified packet core data plane.
 *
 * @summary Gets information about the specified packet core data plane.
 * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2024-02-01/examples/PacketCoreDataPlaneGet.json
 */
async function getPacketCoreDataPlane() {
  const subscriptionId =
    process.env["MOBILENETWORK_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["MOBILENETWORK_RESOURCE_GROUP"] || "rg1";
  const packetCoreControlPlaneName = "testPacketCoreCP";
  const packetCoreDataPlaneName = "testPacketCoreDP";
  const credential = new DefaultAzureCredential();
  const client = new MobileNetworkManagementClient(credential, subscriptionId);
  const result = await client.packetCoreDataPlanes.get(
    resourceGroupName,
    packetCoreControlPlaneName,
    packetCoreDataPlaneName,
  );
  console.log(result);
}
