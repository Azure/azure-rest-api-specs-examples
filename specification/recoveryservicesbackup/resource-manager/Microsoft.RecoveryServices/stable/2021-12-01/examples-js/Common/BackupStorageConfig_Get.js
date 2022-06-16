const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");

async function getVaultStorageConfiguration() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const vaultName = "PySDKBackupTestRsVault";
  const resourceGroupName = "PythonSDKBackupTestRg";
  const credential = new DefaultAzureCredential();
  const client = new RecoveryServicesBackupClient(credential, subscriptionId);
  const result = await client.backupResourceStorageConfigsNonCRR.get(vaultName, resourceGroupName);
  console.log(result);
}

getVaultStorageConfiguration().catch(console.error);
