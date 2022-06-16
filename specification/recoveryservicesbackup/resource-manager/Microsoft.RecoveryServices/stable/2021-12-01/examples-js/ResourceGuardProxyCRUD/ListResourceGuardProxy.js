const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");

async function getVaultGuardProxies() {
  const subscriptionId = "0b352192-dcac-4cc7-992e-a96190ccc68c";
  const vaultName = "sampleVault";
  const resourceGroupName = "SampleResourceGroup";
  const credential = new DefaultAzureCredential();
  const client = new RecoveryServicesBackupClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.resourceGuardProxies.list(vaultName, resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

getVaultGuardProxies().catch(console.error);
