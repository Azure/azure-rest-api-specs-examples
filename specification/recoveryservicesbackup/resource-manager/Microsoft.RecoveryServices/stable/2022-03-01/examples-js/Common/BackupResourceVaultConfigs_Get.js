const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Fetches resource vault config.
 *
 * @summary Fetches resource vault config.
 * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2022-03-01/examples/Common/BackupResourceVaultConfigs_Get.json
 */
async function getVaultSecurityConfig() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const vaultName = "SwaggerTest";
  const resourceGroupName = "SwaggerTestRg";
  const credential = new DefaultAzureCredential();
  const client = new RecoveryServicesBackupClient(credential, subscriptionId);
  const result = await client.backupResourceVaultConfigs.get(vaultName, resourceGroupName);
  console.log(result);
}

getVaultSecurityConfig().catch(console.error);
