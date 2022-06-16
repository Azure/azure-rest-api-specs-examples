const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");

async function getProtectedItemsUsagesSummary() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const vaultName = "testVault";
  const resourceGroupName = "testRG";
  const filter = "type eq 'BackupProtectedItemCountSummary'";
  const options = { filter: filter };
  const credential = new DefaultAzureCredential();
  const client = new RecoveryServicesBackupClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.backupUsageSummaries.list(vaultName, resourceGroupName, options)) {
    resArray.push(item);
  }
  console.log(resArray);
}

getProtectedItemsUsagesSummary().catch(console.error);
