const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Stops a specified packet capture session.
 *
 * @summary Stops a specified packet capture session.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2023-04-01/examples/NetworkWatcherPacketCaptureStop.json
 */
async function stopPacketCapture() {
  const subscriptionId = process.env["NETWORK_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["NETWORK_RESOURCE_GROUP"] || "rg1";
  const networkWatcherName = "nw1";
  const packetCaptureName = "pc1";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.packetCaptures.beginStopAndWait(
    resourceGroupName,
    networkWatcherName,
    packetCaptureName
  );
  console.log(result);
}
