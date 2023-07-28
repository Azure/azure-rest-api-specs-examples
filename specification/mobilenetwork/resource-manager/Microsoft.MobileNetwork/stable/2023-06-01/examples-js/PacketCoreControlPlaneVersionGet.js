const { MobileNetworkManagementClient } = require("@azure/arm-mobilenetwork");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets information about the specified packet core control plane version.
 *
 * @summary Gets information about the specified packet core control plane version.
 * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2023-06-01/examples/PacketCoreControlPlaneVersionGet.json
 */
async function getPacketCoreControlPlaneVersion() {
  const versionName = "PMN-4-11-1";
  const credential = new DefaultAzureCredential();
  const client = new MobileNetworkManagementClient(credential);
  const result = await client.packetCoreControlPlaneVersions.get(versionName);
  console.log(result);
}
