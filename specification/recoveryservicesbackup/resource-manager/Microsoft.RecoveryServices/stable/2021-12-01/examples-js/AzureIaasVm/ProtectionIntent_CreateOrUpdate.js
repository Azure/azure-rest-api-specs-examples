const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");

async function createOrUpdateAzureVMProtectionIntent() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const vaultName = "myVault";
  const resourceGroupName = "myRG";
  const fabricName = "Azure";
  const intentObjectName = "vm;iaasvmcontainerv2;chamsrgtest;chamscandel";
  const parameters = {
    properties: {
      policyId:
        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.RecoveryServices/vaults/myVault/backupPolicies/myPolicy",
      protectionIntentItemType: "AzureResourceItem",
      sourceResourceId:
        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/chamsrgtest/providers/Microsoft.Compute/virtualMachines/chamscandel",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new RecoveryServicesBackupClient(credential, subscriptionId);
  const result = await client.protectionIntentOperations.createOrUpdate(
    vaultName,
    resourceGroupName,
    fabricName,
    intentObjectName,
    parameters
  );
  console.log(result);
}

createOrUpdateAzureVMProtectionIntent().catch(console.error);
