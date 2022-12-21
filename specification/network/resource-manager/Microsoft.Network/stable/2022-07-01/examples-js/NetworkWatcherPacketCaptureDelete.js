const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the specified packet capture session.
 *
 * @summary Deletes the specified packet capture session.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-07-01/examples/NetworkWatcherPacketCaptureDelete.json
 */
async function deletePacketCapture() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const networkWatcherName = "nw1";
  const packetCaptureName = "pc1";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.packetCaptures.beginDeleteAndWait(
    resourceGroupName,
    networkWatcherName,
    packetCaptureName
  );
  console.log(result);
}

deletePacketCapture().catch(console.error);
