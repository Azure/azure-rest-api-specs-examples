const { NetAppManagementClient } = require("@azure/arm-netapp");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the status of the replication
 *
 * @summary Get the status of the replication
 * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2021-10-01/examples/Volumes_ReplicationStatus.json
 */
async function volumesReplicationStatus() {
  const subscriptionId = "D633CC2E-722B-4AE1-B636-BBD9E4C60ED9";
  const resourceGroupName = "myRG";
  const accountName = "account1";
  const poolName = "pool1";
  const volumeName = "volume1";
  const credential = new DefaultAzureCredential();
  const client = new NetAppManagementClient(credential, subscriptionId);
  const result = await client.volumes.replicationStatus(
    resourceGroupName,
    accountName,
    poolName,
    volumeName
  );
  console.log(result);
}

volumesReplicationStatus().catch(console.error);
