const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Cancels a job. This is an asynchronous operation. To know the status of the cancellation, call
GetCancelOperationResult API.
 *
 * @summary Cancels a job. This is an asynchronous operation. To know the status of the cancellation, call
GetCancelOperationResult API.
 * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2022-03-01/examples/Common/TriggerCancelJob.json
 */
async function cancelJob() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const vaultName = "NetSDKTestRsVault";
  const resourceGroupName = "SwaggerTestRg";
  const jobName = "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new RecoveryServicesBackupClient(credential, subscriptionId);
  const result = await client.jobCancellations.trigger(vaultName, resourceGroupName, jobName);
  console.log(result);
}

cancelJob().catch(console.error);
