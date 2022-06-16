const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");

async function deleteProtectionFromAzureVirtualMachine() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const vaultName = "PySDKBackupTestRsVault";
  const resourceGroupName = "PythonSDKBackupTestRg";
  const fabricName = "Azure";
  const containerName = "iaasvmcontainer;iaasvmcontainerv2;pysdktestrg;pysdktestv2vm1";
  const protectedItemName = "vm;iaasvmcontainerv2;pysdktestrg;pysdktestv2vm1";
  const credential = new DefaultAzureCredential();
  const client = new RecoveryServicesBackupClient(credential, subscriptionId);
  const result = await client.protectedItems.delete(
    vaultName,
    resourceGroupName,
    fabricName,
    containerName,
    protectedItemName
  );
  console.log(result);
}

deleteProtectionFromAzureVirtualMachine().catch(console.error);
