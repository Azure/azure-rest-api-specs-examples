const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Provides the result of the refresh operation triggered by the BeginRefresh operation.
 *
 * @summary Provides the result of the refresh operation triggered by the BeginRefresh operation.
 * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2022-03-01/examples/Common/RefreshContainers_OperationResults.json
 */
async function azureVMDiscoveryOperationResult() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const vaultName = "NetSDKTestRsVault";
  const resourceGroupName = "SwaggerTestRg";
  const fabricName = "Azure";
  const operationId = "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new RecoveryServicesBackupClient(credential, subscriptionId);
  const result = await client.protectionContainerRefreshOperationResults.get(
    vaultName,
    resourceGroupName,
    fabricName,
    operationId
  );
  console.log(result);
}

azureVMDiscoveryOperationResult().catch(console.error);
