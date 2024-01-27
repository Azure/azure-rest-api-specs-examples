const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates vault security config.
 *
 * @summary Updates vault security config.
 * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2023-06-01/examples/Common/BackupResourceVaultConfigs_Patch.json
 */
async function updateVaultSecurityConfig() {
  const subscriptionId =
    process.env["RECOVERYSERVICESBACKUP_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const vaultName = "SwaggerTest";
  const resourceGroupName = process.env["RECOVERYSERVICESBACKUP_RESOURCE_GROUP"] || "SwaggerTestRg";
  const parameters = {
    properties: { enhancedSecurityState: "Enabled" },
  };
  const credential = new DefaultAzureCredential();
  const client = new RecoveryServicesBackupClient(credential, subscriptionId);
  const result = await client.backupResourceVaultConfigs.update(
    vaultName,
    resourceGroupName,
    parameters,
  );
  console.log(result);
}
