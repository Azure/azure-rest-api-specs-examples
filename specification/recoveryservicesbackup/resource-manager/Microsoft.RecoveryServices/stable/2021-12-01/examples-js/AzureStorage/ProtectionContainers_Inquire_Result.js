const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");

async function getAzureStorageProtectionContainerOperationResult() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const vaultName = "testvault";
  const resourceGroupName = "test-rg";
  const fabricName = "Azure";
  const containerName = "VMAppContainer;Compute;testRG;testSQL";
  const operationId = "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new RecoveryServicesBackupClient(credential, subscriptionId);
  const result = await client.protectionContainerOperationResults.get(
    vaultName,
    resourceGroupName,
    fabricName,
    containerName,
    operationId
  );
  console.log(result);
}

getAzureStorageProtectionContainerOperationResult().catch(console.error);
