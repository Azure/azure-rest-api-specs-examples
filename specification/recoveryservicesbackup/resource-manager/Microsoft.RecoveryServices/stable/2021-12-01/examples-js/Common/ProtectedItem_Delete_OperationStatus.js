const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");

async function getProtectedItemDeleteOperationStatus() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const vaultName = "PySDKBackupTestRsVault";
  const resourceGroupName = "PythonSDKBackupTestRg";
  const operationId = "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new RecoveryServicesBackupClient(credential, subscriptionId);
  const result = await client.backupOperationStatuses.get(
    vaultName,
    resourceGroupName,
    operationId
  );
  console.log(result);
}

getProtectedItemDeleteOperationStatus().catch(console.error);
