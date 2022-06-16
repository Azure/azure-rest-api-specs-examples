const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");

async function unregisterProtectionContainer() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const vaultName = "testVault";
  const resourceGroupName = "testRg";
  const fabricName = "Azure";
  const containerName = "storagecontainer;Storage;test-rg;teststorage";
  const credential = new DefaultAzureCredential();
  const client = new RecoveryServicesBackupClient(credential, subscriptionId);
  const result = await client.protectionContainers.unregister(
    vaultName,
    resourceGroupName,
    fabricName,
    containerName
  );
  console.log(result);
}

unregisterProtectionContainer().catch(console.error);
