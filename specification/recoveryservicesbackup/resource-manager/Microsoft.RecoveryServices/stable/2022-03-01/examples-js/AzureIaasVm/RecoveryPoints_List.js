const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the backup copies for the backed up item.
 *
 * @summary Lists the backup copies for the backed up item.
 * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2022-03-01/examples/AzureIaasVm/RecoveryPoints_List.json
 */
async function getProtectedAzureVMRecoveryPoints() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const vaultName = "rshvault";
  const resourceGroupName = "rshhtestmdvmrg";
  const fabricName = "Azure";
  const containerName = "IaasVMContainer;iaasvmcontainerv2;rshhtestmdvmrg;rshmdvmsmall";
  const protectedItemName = "VM;iaasvmcontainerv2;rshhtestmdvmrg;rshmdvmsmall";
  const credential = new DefaultAzureCredential();
  const client = new RecoveryServicesBackupClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.recoveryPoints.list(
    vaultName,
    resourceGroupName,
    fabricName,
    containerName,
    protectedItemName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

getProtectedAzureVMRecoveryPoints().catch(console.error);
