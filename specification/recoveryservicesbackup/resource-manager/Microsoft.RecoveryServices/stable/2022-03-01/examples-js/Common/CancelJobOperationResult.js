const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Fetches the result of any operation.
 *
 * @summary Fetches the result of any operation.
 * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2022-03-01/examples/Common/CancelJobOperationResult.json
 */
async function cancelJobOperationResult() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const vaultName = "NetSDKTestRsVault";
  const resourceGroupName = "SwaggerTestRg";
  const jobName = "00000000-0000-0000-0000-000000000000";
  const operationId = "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new RecoveryServicesBackupClient(credential, subscriptionId);
  const result = await client.jobOperationResults.get(
    vaultName,
    resourceGroupName,
    jobName,
    operationId
  );
  console.log(result);
}

cancelJobOperationResult().catch(console.error);
