const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Provides the details of the tiering related sizes and cost.
Status of the operation can be fetched using GetTieringCostOperationStatus API and result using GetTieringCostOperationResult API.
 *
 * @summary Provides the details of the tiering related sizes and cost.
Status of the operation can be fetched using GetTieringCostOperationStatus API and result using GetTieringCostOperationResult API.
 * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2024-04-01/examples/TieringCost/FetchTieringCostForProtectedItem.json
 */
async function getTheTieringSavingsCostInfoForProtectedItem() {
  const subscriptionId =
    process.env["RECOVERYSERVICESBACKUP_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["RECOVERYSERVICESBACKUP_RESOURCE_GROUP"] || "netsdktestrg";
  const vaultName = "testVault";
  const parameters = {
    containerName: "IaasVMContainer;iaasvmcontainerv2;netsdktestrg;netvmtestv2vm1",
    objectType: "FetchTieringCostSavingsInfoForProtectedItemRequest",
    protectedItemName: "VM;iaasvmcontainerv2;netsdktestrg;netvmtestv2vm1",
    sourceTierType: "HardenedRP",
    targetTierType: "ArchivedRP",
  };
  const credential = new DefaultAzureCredential();
  const client = new RecoveryServicesBackupClient(credential, subscriptionId);
  const result = await client.fetchTieringCost.beginPostAndWait(
    resourceGroupName,
    vaultName,
    parameters,
  );
  console.log(result);
}
