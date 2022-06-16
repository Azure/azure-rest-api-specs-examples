const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");

async function prepareDataMove() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const vaultName = "source-rsv";
  const resourceGroupName = "sourceRG";
  const parameters = {
    dataMoveLevel: "Vault",
    targetRegion: "USGov Virginia",
    targetResourceId:
      "/subscriptions/04cf684a-d41f-4550-9f70-7708a3a2283b/resourceGroups/targetRG/providers/Microsoft.RecoveryServices/vaults/target-rsv",
  };
  const credential = new DefaultAzureCredential();
  const client = new RecoveryServicesBackupClient(credential, subscriptionId);
  const result = await client.beginBMSPrepareDataMoveAndWait(
    vaultName,
    resourceGroupName,
    parameters
  );
  console.log(result);
}

prepareDataMove().catch(console.error);
