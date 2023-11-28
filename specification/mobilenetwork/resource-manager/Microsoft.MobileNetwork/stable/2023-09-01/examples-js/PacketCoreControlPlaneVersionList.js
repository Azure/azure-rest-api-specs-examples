const { MobileNetworkManagementClient } = require("@azure/arm-mobilenetwork");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all supported packet core control planes versions.
 *
 * @summary Lists all supported packet core control planes versions.
 * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2023-09-01/examples/PacketCoreControlPlaneVersionList.json
 */
async function getSupportedPacketCoreControlPlaneVersions() {
  const credential = new DefaultAzureCredential();
  const client = new MobileNetworkManagementClient(credential);
  const resArray = new Array();
  for await (let item of client.packetCoreControlPlaneVersions.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}
