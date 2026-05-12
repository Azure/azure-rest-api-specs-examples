const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to provides the details of the tiering related sizes and cost.
 * Status of the operation can be fetched using GetTieringCostOperationStatus API and result using GetTieringCostOperationResult API.
 *
 * @summary provides the details of the tiering related sizes and cost.
 * Status of the operation can be fetched using GetTieringCostOperationStatus API and result using GetTieringCostOperationResult API.
 * x-ms-original-file: 2026-01-31-preview/TieringCost/FetchTieringCostForVault.json
 */
async function getTheTieringSavingsCostInfoForVault() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new RecoveryServicesBackupClient(credential, subscriptionId);
  const result = await client.fetchTieringCost.post("netsdktestrg", "testVault", {
    objectType: "FetchTieringCostSavingsInfoForVaultRequest",
    sourceTierType: "HardenedRP",
    targetTierType: "ArchivedRP",
  });
  console.log(result);
}
