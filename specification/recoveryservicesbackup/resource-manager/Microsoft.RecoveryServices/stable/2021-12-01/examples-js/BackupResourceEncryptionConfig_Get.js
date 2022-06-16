const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");

async function getVaultEncryptionConfiguration() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const vaultName = "rishTestVault";
  const resourceGroupName = "rishgrp";
  const credential = new DefaultAzureCredential();
  const client = new RecoveryServicesBackupClient(credential, subscriptionId);
  const result = await client.backupResourceEncryptionConfigs.get(vaultName, resourceGroupName);
  console.log(result);
}

getVaultEncryptionConfiguration().catch(console.error);
