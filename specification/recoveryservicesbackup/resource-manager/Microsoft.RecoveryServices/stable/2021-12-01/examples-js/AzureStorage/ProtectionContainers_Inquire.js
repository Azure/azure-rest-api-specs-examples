const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");

async function inquireAzureStorageProtectionContainers() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const vaultName = "testvault";
  const resourceGroupName = "test-rg";
  const fabricName = "Azure";
  const containerName = "storagecontainer;Storage;test-rg;teststorage";
  const credential = new DefaultAzureCredential();
  const client = new RecoveryServicesBackupClient(credential, subscriptionId);
  const result = await client.protectionContainers.inquire(
    vaultName,
    resourceGroupName,
    fabricName,
    containerName
  );
  console.log(result);
}

inquireAzureStorageProtectionContainers().catch(console.error);
