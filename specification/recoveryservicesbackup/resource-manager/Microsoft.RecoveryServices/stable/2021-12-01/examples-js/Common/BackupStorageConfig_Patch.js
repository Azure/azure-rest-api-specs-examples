const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");

async function updateVaultStorageConfiguration() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const vaultName = "PySDKBackupTestRsVault";
  const resourceGroupName = "PythonSDKBackupTestRg";
  const parameters = {
    properties: {
      storageType: "LocallyRedundant",
      storageTypeState: "Unlocked",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new RecoveryServicesBackupClient(credential, subscriptionId);
  const result = await client.backupResourceStorageConfigsNonCRR.patch(
    vaultName,
    resourceGroupName,
    parameters
  );
  console.log(result);
}

updateVaultStorageConfiguration().catch(console.error);
