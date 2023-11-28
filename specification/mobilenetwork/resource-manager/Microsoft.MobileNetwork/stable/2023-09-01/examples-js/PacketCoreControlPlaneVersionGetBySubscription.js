const { MobileNetworkManagementClient } = require("@azure/arm-mobilenetwork");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets information about the specified packet core control plane version.
 *
 * @summary Gets information about the specified packet core control plane version.
 * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2023-09-01/examples/PacketCoreControlPlaneVersionGetBySubscription.json
 */
async function getPacketCoreControlPlaneVersionBySubscription() {
  const subscriptionId =
    process.env["MOBILENETWORK_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const versionName = "PMN-4-11-1";
  const credential = new DefaultAzureCredential();
  const client = new MobileNetworkManagementClient(credential, subscriptionId);
  const result = await client.packetCoreControlPlaneVersions.getBySubscription(versionName);
  console.log(result);
}
