const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");

async function deleteProtectionIntentFromItem() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const vaultName = "myVault";
  const resourceGroupName = "myRG";
  const fabricName = "Azure";
  const intentObjectName = "249D9B07-D2EF-4202-AA64-65F35418564E";
  const credential = new DefaultAzureCredential();
  const client = new RecoveryServicesBackupClient(credential, subscriptionId);
  const result = await client.protectionIntentOperations.delete(
    vaultName,
    resourceGroupName,
    fabricName,
    intentObjectName
  );
  console.log(result);
}

deleteProtectionIntentFromItem().catch(console.error);
