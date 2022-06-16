const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");

async function listBackupProtectionContainers() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const vaultName = "testVault";
  const resourceGroupName = "testRg";
  const filter = "backupManagementType eq 'AzureWorkload'";
  const options = { filter: filter };
  const credential = new DefaultAzureCredential();
  const client = new RecoveryServicesBackupClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.backupProtectionContainers.list(
    vaultName,
    resourceGroupName,
    options
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

listBackupProtectionContainers().catch(console.error);
