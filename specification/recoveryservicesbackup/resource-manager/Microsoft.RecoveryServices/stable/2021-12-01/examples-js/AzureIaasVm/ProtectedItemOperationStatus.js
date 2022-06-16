const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");

async function getOperationStatusOfProtectedVM() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const vaultName = "NetSDKTestRsVault";
  const resourceGroupName = "SwaggerTestRg";
  const fabricName = "Azure";
  const containerName = "IaasVMContainer;iaasvmcontainerv2;netsdktestrg;netvmtestv2vm1";
  const protectedItemName = "VM;iaasvmcontainerv2;netsdktestrg;netvmtestv2vm1";
  const operationId = "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new RecoveryServicesBackupClient(credential, subscriptionId);
  const result = await client.protectedItemOperationStatuses.get(
    vaultName,
    resourceGroupName,
    fabricName,
    containerName,
    protectedItemName,
    operationId
  );
  console.log(result);
}

getOperationStatusOfProtectedVM().catch(console.error);
