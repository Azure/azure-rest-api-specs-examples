const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to provides the information of the backed up data identified using RecoveryPointID. This is an asynchronous operation.
 * To know the status of the operation, call the GetProtectedItemOperationResult API.
 *
 * @summary provides the information of the backed up data identified using RecoveryPointID. This is an asynchronous operation.
 * To know the status of the operation, call the GetProtectedItemOperationResult API.
 * x-ms-original-file: 2026-10-01/AzureWorkload/RecoveryPoints_Get_Snapshot.json
 */
async function getAzureWorkloadSQLSnapshotRecoveryPointDetails() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new RecoveryServicesBackupClient(credential, subscriptionId);
  const result = await client.recoveryPoints.get(
    "testVault",
    "testRG",
    "Azure",
    "VMAppContainer;Compute;testRG;sqlVm",
    "SQLDataBase;mssqlserver;inventory",
    "1700000000000",
  );
  console.log(result);
}
