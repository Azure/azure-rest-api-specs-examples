const { NetAppManagementClient } = require("@azure/arm-netapp");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Patch the specified NetApp Backup Vault
 *
 * @summary Patch the specified NetApp Backup Vault
 * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2023-11-01/examples/BackupVaults_Update.json
 */
async function backupVaultsUpdate() {
  const subscriptionId =
    process.env["NETAPP_SUBSCRIPTION_ID"] || "D633CC2E-722B-4AE1-B636-BBD9E4C60ED9";
  const resourceGroupName = process.env["NETAPP_RESOURCE_GROUP"] || "myRG";
  const accountName = "account1";
  const backupVaultName = "backupVault1";
  const body = { tags: { tag1: "Value1" } };
  const credential = new DefaultAzureCredential();
  const client = new NetAppManagementClient(credential, subscriptionId);
  const result = await client.backupVaults.beginUpdateAndWait(
    resourceGroupName,
    accountName,
    backupVaultName,
    body,
  );
  console.log(result);
}
