const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");

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
