const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");

async function getProtectionContainerDetails() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const vaultName = "testVault";
  const resourceGroupName = "testRg";
  const fabricName = "Azure";
  const containerName = "VMAppContainer;Compute;testRG;testSQL";
  const credential = new DefaultAzureCredential();
  const client = new RecoveryServicesBackupClient(credential, subscriptionId);
  const result = await client.protectionContainers.get(
    vaultName,
    resourceGroupName,
    fabricName,
    containerName
  );
  console.log(result);
}

getProtectionContainerDetails().catch(console.error);
