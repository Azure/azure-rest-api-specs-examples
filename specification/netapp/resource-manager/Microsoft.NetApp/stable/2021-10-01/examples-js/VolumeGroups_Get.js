const { NetAppManagementClient } = require("@azure/arm-netapp");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get details of the specified volume group
 *
 * @summary Get details of the specified volume group
 * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2021-10-01/examples/VolumeGroups_Get.json
 */
async function volumeGroupsGet() {
  const subscriptionId = "D633CC2E-722B-4AE1-B636-BBD9E4C60ED9";
  const resourceGroupName = "myRG";
  const accountName = "account1";
  const volumeGroupName = "group1";
  const credential = new DefaultAzureCredential();
  const client = new NetAppManagementClient(credential, subscriptionId);
  const result = await client.volumeGroups.get(resourceGroupName, accountName, volumeGroupName);
  console.log(result);
}

volumeGroupsGet().catch(console.error);
