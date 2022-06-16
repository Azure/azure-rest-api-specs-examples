const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates vault security config.
 *
 * @summary Updates vault security config.
 * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2021-12-01/examples/Common/BackupResourceVaultConfigs_Patch.json
 */
async function updateVaultSecurityConfig() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const vaultName = "SwaggerTest";
  const resourceGroupName = "SwaggerTestRg";
  const parameters = {
    properties: { enhancedSecurityState: "Enabled" },
  };
  const credential = new DefaultAzureCredential();
  const client = new RecoveryServicesBackupClient(credential, subscriptionId);
  const result = await client.backupResourceVaultConfigs.update(
    vaultName,
    resourceGroupName,
    parameters
  );
  console.log(result);
}

updateVaultSecurityConfig().catch(console.error);
