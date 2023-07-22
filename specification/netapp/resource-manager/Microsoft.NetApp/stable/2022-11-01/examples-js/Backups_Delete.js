const { NetAppManagementClient } = require("@azure/arm-netapp");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete a backup of the volume
 *
 * @summary Delete a backup of the volume
 * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2022-11-01/examples/Backups_Delete.json
 */
async function backupsDelete() {
  const subscriptionId =
    process.env["NETAPP_SUBSCRIPTION_ID"] || "D633CC2E-722B-4AE1-B636-BBD9E4C60ED9";
  const resourceGroupName = process.env["NETAPP_RESOURCE_GROUP"] || "resourceGroup";
  const accountName = "accountName";
  const poolName = "poolName";
  const volumeName = "volumeName";
  const backupName = "backupName";
  const credential = new DefaultAzureCredential();
  const client = new NetAppManagementClient(credential, subscriptionId);
  const result = await client.backups.beginDeleteAndWait(
    resourceGroupName,
    accountName,
    poolName,
    volumeName,
    backupName
  );
  console.log(result);
}
