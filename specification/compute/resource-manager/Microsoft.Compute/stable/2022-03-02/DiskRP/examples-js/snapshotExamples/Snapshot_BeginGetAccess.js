const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Grants access to a snapshot.
 *
 * @summary Grants access to a snapshot.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-02/DiskRP/examples/snapshotExamples/Snapshot_BeginGetAccess.json
 */
async function getASasOnASnapshot() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const snapshotName = "mySnapshot";
  const grantAccessData = {
    access: "Read",
    durationInSeconds: 300,
  };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.snapshots.beginGrantAccessAndWait(
    resourceGroupName,
    snapshotName,
    grantAccessData
  );
  console.log(result);
}

getASasOnASnapshot().catch(console.error);
