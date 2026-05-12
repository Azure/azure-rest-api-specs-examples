const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to provides the details of the tiering related sizes and cost.
 * Status of the operation can be fetched using GetTieringCostOperationStatus API and result using GetTieringCostOperationResult API.
 *
 * @summary provides the details of the tiering related sizes and cost.
 * Status of the operation can be fetched using GetTieringCostOperationStatus API and result using GetTieringCostOperationResult API.
 * x-ms-original-file: 2026-01-31-preview/TieringCost/FetchTieringCostForProtectedItem.json
 */
async function getTheTieringSavingsCostInfoForProtectedItem() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new RecoveryServicesBackupClient(credential, subscriptionId);
  const result = await client.fetchTieringCost.post("netsdktestrg", "testVault", {
    containerName: "IaasVMContainer;iaasvmcontainerv2;netsdktestrg;netvmtestv2vm1",
    objectType: "FetchTieringCostSavingsInfoForProtectedItemRequest",
    protectedItemName: "VM;iaasvmcontainerv2;netsdktestrg;netvmtestv2vm1",
    sourceTierType: "HardenedRP",
    targetTierType: "ArchivedRP",
  });
  console.log(result);
}
