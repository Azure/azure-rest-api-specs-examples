const { NetAppManagementClient } = require("@azure/arm-netapp");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Get details of the specified volume group
 *
 * @summary Get details of the specified volume group
 * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2025-03-01/examples/VolumeGroups_Get_SapHana.json
 */
async function volumeGroupsGetSapHana() {
  const subscriptionId =
    process.env["NETAPP_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["NETAPP_RESOURCE_GROUP"] || "myRG";
  const accountName = "account1";
  const volumeGroupName = "group1";
  const credential = new DefaultAzureCredential();
  const client = new NetAppManagementClient(credential, subscriptionId);
  const result = await client.volumeGroups.get(resourceGroupName, accountName, volumeGroupName);
  console.log(result);
}
