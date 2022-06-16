const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");

async function listWorkloadItemsInContainer() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const vaultName = "suchandr-seacan-rsv";
  const resourceGroupName = "testRg";
  const fabricName = "Azure";
  const containerName = "VMAppContainer;Compute;bvtdtestag;sqlserver-1";
  const filter = "backupManagementType eq 'AzureWorkload'";
  const options = { filter: filter };
  const credential = new DefaultAzureCredential();
  const client = new RecoveryServicesBackupClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.backupWorkloadItems.list(
    vaultName,
    resourceGroupName,
    fabricName,
    containerName,
    options
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

listWorkloadItemsInContainer().catch(console.error);
