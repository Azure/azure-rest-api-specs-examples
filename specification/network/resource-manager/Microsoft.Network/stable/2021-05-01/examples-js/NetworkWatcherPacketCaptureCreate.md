Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-network_27.0.0/sdk/network/arm-network/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create and start a packet capture on the specified VM.
 *
 * @summary Create and start a packet capture on the specified VM.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/NetworkWatcherPacketCaptureCreate.json
 */
async function createPacketCapture() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const networkWatcherName = "nw1";
  const packetCaptureName = "pc1";
  const parameters = {
    bytesToCapturePerPacket: 10000,
    filters: [{ localIPAddress: "10.0.0.4", localPort: "80", protocol: "TCP" }],
    storageLocation: {
      filePath: "D:capturepc1.cap",
      storageId:
        "/subscriptions/subid/resourceGroups/rg2/providers/Microsoft.Storage/storageAccounts/pcstore",
      storagePath: "https://mytestaccountname.blob.core.windows.net/capture/pc1.cap",
    },
    target:
      "/subscriptions/subid/resourceGroups/rg2/providers/Microsoft.Compute/virtualMachines/vm1",
    timeLimitInSeconds: 100,
    totalBytesPerSession: 100000,
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.packetCaptures.beginCreateAndWait(
    resourceGroupName,
    networkWatcherName,
    packetCaptureName,
    parameters
  );
  console.log(result);
}

createPacketCapture().catch(console.error);
```
