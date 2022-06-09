Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-compute_19.0.0/sdk/compute/arm-compute/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates (patches) a snapshot.
 *
 * @summary Updates (patches) a snapshot.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-02/DiskRP/examples/snapshotExamples/Snapshot_Update_WithAcceleratedNetwork.json
 */
async function updateASnapshotWithAcceleratedNetworking() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const snapshotName = "mySnapshot";
  const snapshot = {
    diskSizeGB: 20,
    supportedCapabilities: { acceleratedNetwork: false },
    tags: { department: "Development", project: "UpdateSnapshots" },
  };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.snapshots.beginUpdateAndWait(
    resourceGroupName,
    snapshotName,
    snapshot
  );
  console.log(result);
}

updateASnapshotWithAcceleratedNetworking().catch(console.error);
```
