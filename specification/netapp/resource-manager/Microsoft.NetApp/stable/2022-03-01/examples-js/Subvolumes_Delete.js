const { NetAppManagementClient } = require("@azure/arm-netapp");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete subvolume
 *
 * @summary Delete subvolume
 * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2022-03-01/examples/Subvolumes_Delete.json
 */
async function subvolumesDelete() {
  const subscriptionId = "D633CC2E-722B-4AE1-B636-BBD9E4C60ED9";
  const resourceGroupName = "myRG";
  const accountName = "account1";
  const poolName = "pool1";
  const volumeName = "volume1";
  const subvolumeName = "subvolume1";
  const credential = new DefaultAzureCredential();
  const client = new NetAppManagementClient(credential, subscriptionId);
  const result = await client.subvolumes.beginDeleteAndWait(
    resourceGroupName,
    accountName,
    poolName,
    volumeName,
    subvolumeName
  );
  console.log(result);
}

subvolumesDelete().catch(console.error);
