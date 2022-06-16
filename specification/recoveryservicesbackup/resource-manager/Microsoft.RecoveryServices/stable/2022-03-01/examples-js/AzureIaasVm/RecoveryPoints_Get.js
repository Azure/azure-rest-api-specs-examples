const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Provides the information of the backed up data identified using RecoveryPointID. This is an asynchronous operation.
To know the status of the operation, call the GetProtectedItemOperationResult API.
 *
 * @summary Provides the information of the backed up data identified using RecoveryPointID. This is an asynchronous operation.
To know the status of the operation, call the GetProtectedItemOperationResult API.
 * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2022-03-01/examples/AzureIaasVm/RecoveryPoints_Get.json
 */
async function getAzureVMRecoveryPointDetails() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const vaultName = "rshvault";
  const resourceGroupName = "rshhtestmdvmrg";
  const fabricName = "Azure";
  const containerName = "IaasVMContainer;iaasvmcontainerv2;rshhtestmdvmrg;rshmdvmsmall";
  const protectedItemName = "VM;iaasvmcontainerv2;rshhtestmdvmrg;rshmdvmsmall";
  const recoveryPointId = "26083826328862";
  const credential = new DefaultAzureCredential();
  const client = new RecoveryServicesBackupClient(credential, subscriptionId);
  const result = await client.recoveryPoints.get(
    vaultName,
    resourceGroupName,
    fabricName,
    containerName,
    protectedItemName,
    recoveryPointId
  );
  console.log(result);
}

getAzureVMRecoveryPointDetails().catch(console.error);
