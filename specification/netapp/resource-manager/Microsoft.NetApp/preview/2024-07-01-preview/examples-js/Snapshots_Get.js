const { NetAppManagementClient } = require("@azure/arm-netapp");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get details of the specified snapshot
 *
 * @summary Get details of the specified snapshot
 * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/preview/2024-07-01-preview/examples/Snapshots_Get.json
 */
async function snapshotsGet() {
  const subscriptionId =
    process.env["NETAPP_SUBSCRIPTION_ID"] || "D633CC2E-722B-4AE1-B636-BBD9E4C60ED9";
  const resourceGroupName = process.env["NETAPP_RESOURCE_GROUP"] || "myRG";
  const accountName = "account1";
  const poolName = "pool1";
  const volumeName = "volume1";
  const snapshotName = "snapshot1";
  const credential = new DefaultAzureCredential();
  const client = new NetAppManagementClient(credential, subscriptionId);
  const result = await client.snapshots.get(
    resourceGroupName,
    accountName,
    poolName,
    volumeName,
    snapshotName,
  );
  console.log(result);
}
