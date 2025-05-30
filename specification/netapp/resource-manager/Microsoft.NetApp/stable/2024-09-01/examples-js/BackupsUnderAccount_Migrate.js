const { NetAppManagementClient } = require("@azure/arm-netapp");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Migrate the backups under a NetApp account to backup vault
 *
 * @summary Migrate the backups under a NetApp account to backup vault
 * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2024-09-01/examples/BackupsUnderAccount_Migrate.json
 */
async function backupsUnderAccountMigrate() {
  const subscriptionId =
    process.env["NETAPP_SUBSCRIPTION_ID"] || "D633CC2E-722B-4AE1-B636-BBD9E4C60ED9";
  const resourceGroupName = process.env["NETAPP_RESOURCE_GROUP"] || "myRG";
  const accountName = "account1";
  const body = {
    backupVaultId:
      "/subscriptions/D633CC2E-722B-4AE1-B636-BBD9E4C60ED9/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/backupVaults/backupVault1",
  };
  const credential = new DefaultAzureCredential();
  const client = new NetAppManagementClient(credential, subscriptionId);
  const result = await client.backupsUnderAccount.beginMigrateBackupsAndWait(
    resourceGroupName,
    accountName,
    body,
  );
  console.log(result);
}
