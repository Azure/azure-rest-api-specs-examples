const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Move recovery point from one datastore to another store.
 *
 * @summary Move recovery point from one datastore to another store.
 * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2022-03-01/examples/TriggerRecoveryPointMove_Post.json
 */
async function triggerRpMoveOperation() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const vaultName = "testVault";
  const resourceGroupName = "netsdktestrg";
  const fabricName = "Azure";
  const containerName = "IaasVMContainer;iaasvmcontainerv2;netsdktestrg;netvmtestv2vm1";
  const protectedItemName = "VM;iaasvmcontainerv2;netsdktestrg;netvmtestv2vm1";
  const recoveryPointId = "348916168024334";
  const parameters = {
    objectType: "MoveRPAcrossTiersRequest",
    sourceTierType: "HardenedRP",
    targetTierType: "ArchivedRP",
  };
  const credential = new DefaultAzureCredential();
  const client = new RecoveryServicesBackupClient(credential, subscriptionId);
  const result = await client.beginMoveRecoveryPointAndWait(
    vaultName,
    resourceGroupName,
    fabricName,
    containerName,
    protectedItemName,
    recoveryPointId,
    parameters
  );
  console.log(result);
}

triggerRpMoveOperation().catch(console.error);
