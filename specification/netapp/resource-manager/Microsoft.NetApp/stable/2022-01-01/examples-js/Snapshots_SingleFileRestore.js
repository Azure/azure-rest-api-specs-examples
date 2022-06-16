const { NetAppManagementClient } = require("@azure/arm-netapp");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Restore the specified files from the specified snapshot to the active filesystem
 *
 * @summary Restore the specified files from the specified snapshot to the active filesystem
 * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2022-01-01/examples/Snapshots_SingleFileRestore.json
 */
async function snapshotsSingleFileRestore() {
  const subscriptionId = "D633CC2E-722B-4AE1-B636-BBD9E4C60ED9";
  const resourceGroupName = "myRG";
  const accountName = "account1";
  const poolName = "pool1";
  const volumeName = "volume1";
  const snapshotName = "snapshot1";
  const body = {
    filePaths: ["/dir1/customer1.db", "/dir1/customer2.db"],
  };
  const credential = new DefaultAzureCredential();
  const client = new NetAppManagementClient(credential, subscriptionId);
  const result = await client.snapshots.beginRestoreFilesAndWait(
    resourceGroupName,
    accountName,
    poolName,
    volumeName,
    snapshotName,
    body
  );
  console.log(result);
}

snapshotsSingleFileRestore().catch(console.error);
