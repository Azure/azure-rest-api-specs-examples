const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");

async function triggerBackup() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const vaultName = "linuxRsVault";
  const resourceGroupName = "linuxRsVaultRG";
  const fabricName = "Azure";
  const containerName = "IaasVMContainer;iaasvmcontainerv2;testrg;v1win2012r";
  const protectedItemName = "VM;iaasvmcontainerv2;testrg;v1win2012r";
  const parameters = {
    properties: { objectType: "IaasVMBackupRequest" },
  };
  const credential = new DefaultAzureCredential();
  const client = new RecoveryServicesBackupClient(credential, subscriptionId);
  const result = await client.backups.trigger(
    vaultName,
    resourceGroupName,
    fabricName,
    containerName,
    protectedItemName,
    parameters
  );
  console.log(result);
}

triggerBackup().catch(console.error);
