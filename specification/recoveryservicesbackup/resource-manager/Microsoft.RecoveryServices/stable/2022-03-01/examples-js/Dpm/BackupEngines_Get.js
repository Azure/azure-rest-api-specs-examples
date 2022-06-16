const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns backup management server registered to Recovery Services Vault.
 *
 * @summary Returns backup management server registered to Recovery Services Vault.
 * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2022-03-01/examples/Dpm/BackupEngines_Get.json
 */
async function getDpmOrAzureBackupServerOrLajollaBackupEngineDetails() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const vaultName = "testVault";
  const resourceGroupName = "testRG";
  const backupEngineName = "testServer";
  const credential = new DefaultAzureCredential();
  const client = new RecoveryServicesBackupClient(credential, subscriptionId);
  const result = await client.backupEngines.get(vaultName, resourceGroupName, backupEngineName);
  console.log(result);
}

getDpmOrAzureBackupServerOrLajollaBackupEngineDetails().catch(console.error);
