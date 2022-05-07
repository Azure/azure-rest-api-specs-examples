Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-storagesync_9.0.1/sdk/storagesync/arm-storagesync/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { MicrosoftStorageSync } = require("@azure/arm-storagesync");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Post Restore a given CloudEndpoint.
 *
 * @summary Post Restore a given CloudEndpoint.
 * x-ms-original-file: specification/storagesync/resource-manager/Microsoft.StorageSync/stable/2020-09-01/examples/CloudEndpoints_PostRestore.json
 */
async function cloudEndpointsPostRestore() {
  const subscriptionId = "52b8da2f-61e0-4a1f-8dde-336911f367fb";
  const resourceGroupName = "SampleResourceGroup_1";
  const storageSyncServiceName = "SampleStorageSyncService_1";
  const syncGroupName = "SampleSyncGroup_1";
  const cloudEndpointName = "SampleCloudEndpoint_1";
  const parameters = {
    azureFileShareUri:
      "https://hfsazbackupdevintncus2.file.core.test-cint.azure-test.net/sampleFileShare",
    restoreFileSpec: [
      { path: "text1.txt", isdir: false },
      { path: "MyDir", isdir: true },
      { path: "MyDir/SubDir", isdir: false },
      { path: "MyDir/SubDir/File1.pdf", isdir: false },
    ],
    sourceAzureFileShareUri:
      "https://hfsazbackupdevintncus2.file.core.test-cint.azure-test.net/sampleFileShare",
    status: "Succeeded",
  };
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftStorageSync(credential, subscriptionId);
  const result = await client.cloudEndpoints.beginPostRestoreAndWait(
    resourceGroupName,
    storageSyncServiceName,
    syncGroupName,
    cloudEndpointName,
    parameters
  );
  console.log(result);
}

cloudEndpointsPostRestore().catch(console.error);
```
