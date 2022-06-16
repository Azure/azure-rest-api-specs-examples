const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Fetches operation status for data move operation on vault
 *
 * @summary Fetches operation status for data move operation on vault
 * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2022-03-01/examples/BackupDataMove/BackupDataMoveOperationStatus_Get.json
 */
async function getOperationStatus() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const vaultName = "source-rsv";
  const resourceGroupName = "sourceRG";
  const operationId = "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new RecoveryServicesBackupClient(credential, subscriptionId);
  const result = await client.getOperationStatus(vaultName, resourceGroupName, operationId);
  console.log(result);
}

getOperationStatus().catch(console.error);
