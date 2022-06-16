const { NetAppManagementClient } = require("@azure/arm-netapp");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create the specified snapshot within the given volume
 *
 * @summary Create the specified snapshot within the given volume
 * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2022-01-01/examples/Snapshots_Create.json
 */
async function snapshotsCreate() {
  const subscriptionId = "D633CC2E-722B-4AE1-B636-BBD9E4C60ED9";
  const resourceGroupName = "myRG";
  const accountName = "account1";
  const poolName = "pool1";
  const volumeName = "volume1";
  const snapshotName = "snapshot1";
  const body = { location: "eastus" };
  const credential = new DefaultAzureCredential();
  const client = new NetAppManagementClient(credential, subscriptionId);
  const result = await client.snapshots.beginCreateAndWait(
    resourceGroupName,
    accountName,
    poolName,
    volumeName,
    snapshotName,
    body
  );
  console.log(result);
}

snapshotsCreate().catch(console.error);
