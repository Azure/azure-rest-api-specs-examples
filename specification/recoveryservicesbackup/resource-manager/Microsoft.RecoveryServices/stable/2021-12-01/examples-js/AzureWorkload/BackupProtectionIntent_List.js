const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");

async function listProtectionIntentWithBackupManagementTypeFilter() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const vaultName = "myVault";
  const resourceGroupName = "myRG";
  const credential = new DefaultAzureCredential();
  const client = new RecoveryServicesBackupClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.backupProtectionIntent.list(vaultName, resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listProtectionIntentWithBackupManagementTypeFilter().catch(console.error);
