const { NetAppManagementClient } = require("@azure/arm-netapp");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete the specified Backup for a Netapp Account
 *
 * @summary Delete the specified Backup for a Netapp Account
 * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2022-01-01/examples/Backups_Account_Delete.json
 */
async function accountBackupsDelete() {
  const subscriptionId = "D633CC2E-722B-4AE1-B636-BBD9E4C60ED9";
  const resourceGroupName = "resourceGroup";
  const accountName = "accountName";
  const backupName = "backupName";
  const credential = new DefaultAzureCredential();
  const client = new NetAppManagementClient(credential, subscriptionId);
  const result = await client.accountBackups.beginDeleteAndWait(
    resourceGroupName,
    accountName,
    backupName
  );
  console.log(result);
}

accountBackupsDelete().catch(console.error);
