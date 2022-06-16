const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");

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
