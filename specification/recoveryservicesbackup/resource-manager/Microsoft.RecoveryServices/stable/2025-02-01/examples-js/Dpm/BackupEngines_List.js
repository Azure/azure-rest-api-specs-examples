const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Backup management servers registered to Recovery Services Vault. Returns a pageable list of servers.
 *
 * @summary Backup management servers registered to Recovery Services Vault. Returns a pageable list of servers.
 * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2025-02-01/examples/Dpm/BackupEngines_List.json
 */
async function listDpmOrAzureBackupServerOrLajollaBackupEngines() {
  const subscriptionId =
    process.env["RECOVERYSERVICESBACKUP_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const vaultName = "testVault";
  const resourceGroupName = process.env["RECOVERYSERVICESBACKUP_RESOURCE_GROUP"] || "testRG";
  const credential = new DefaultAzureCredential();
  const client = new RecoveryServicesBackupClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.backupEngines.list(vaultName, resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
