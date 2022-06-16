const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");

async function getOperationResultForPrepareDataMove() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const vaultName = "source-rsv";
  const resourceGroupName = "sourceRG";
  const operationId = "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new RecoveryServicesBackupClient(credential, subscriptionId);
  const result = await client.bMSPrepareDataMoveOperationResult.get(
    vaultName,
    resourceGroupName,
    operationId
  );
  console.log(result);
}

getOperationResultForPrepareDataMove().catch(console.error);
