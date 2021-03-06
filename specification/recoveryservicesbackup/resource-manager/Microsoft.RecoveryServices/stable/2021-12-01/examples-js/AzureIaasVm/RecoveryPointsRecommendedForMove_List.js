const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");

async function getProtectedAzureVMRecoveryPointsRecommendedForMove() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const vaultName = "rshvault";
  const resourceGroupName = "rshhtestmdvmrg";
  const fabricName = "Azure";
  const containerName = "IaasVMContainer;iaasvmcontainerv2;rshhtestmdvmrg;rshmdvmsmall";
  const protectedItemName = "VM;iaasvmcontainerv2;rshhtestmdvmrg;rshmdvmsmall";
  const parameters = {
    excludedRPList: ["348916168024334", "348916168024335"],
    objectType: "ListRecoveryPointsRecommendedForMoveRequest",
  };
  const credential = new DefaultAzureCredential();
  const client = new RecoveryServicesBackupClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.recoveryPointsRecommendedForMove.list(
    vaultName,
    resourceGroupName,
    fabricName,
    containerName,
    protectedItemName,
    parameters
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

getProtectedAzureVMRecoveryPointsRecommendedForMove().catch(console.error);
