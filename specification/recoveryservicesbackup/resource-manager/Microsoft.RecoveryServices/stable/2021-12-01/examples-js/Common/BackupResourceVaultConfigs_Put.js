const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");

async function updateVaultSecurityConfig() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const vaultName = "SwaggerTest";
  const resourceGroupName = "SwaggerTestRg";
  const parameters = {
    properties: {
      enhancedSecurityState: "Enabled",
      softDeleteFeatureState: "Enabled",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new RecoveryServicesBackupClient(credential, subscriptionId);
  const result = await client.backupResourceVaultConfigs.put(
    vaultName,
    resourceGroupName,
    parameters
  );
  console.log(result);
}

updateVaultSecurityConfig().catch(console.error);
