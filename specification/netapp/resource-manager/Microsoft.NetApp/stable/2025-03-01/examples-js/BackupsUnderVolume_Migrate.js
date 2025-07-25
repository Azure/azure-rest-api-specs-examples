const { NetAppManagementClient } = require("@azure/arm-netapp");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Migrate the backups under volume to backup vault
 *
 * @summary Migrate the backups under volume to backup vault
 * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2025-03-01/examples/BackupsUnderVolume_Migrate.json
 */
async function backupsUnderVolumeMigrate() {
  const subscriptionId =
    process.env["NETAPP_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["NETAPP_RESOURCE_GROUP"] || "myRG";
  const accountName = "account1";
  const poolName = "pool1";
  const volumeName = "volume1";
  const body = {
    backupVaultId:
      "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/backupVaults/backupVault1",
  };
  const credential = new DefaultAzureCredential();
  const client = new NetAppManagementClient(credential, subscriptionId);
  const result = await client.backupsUnderVolume.beginMigrateBackupsAndWait(
    resourceGroupName,
    accountName,
    poolName,
    volumeName,
    body,
  );
  console.log(result);
}
