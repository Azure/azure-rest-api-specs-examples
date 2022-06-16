const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Triggers export of jobs specified by filters and returns an OperationID to track.
 *
 * @summary Triggers export of jobs specified by filters and returns an OperationID to track.
 * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2022-03-01/examples/Common/TriggerExportJobs.json
 */
async function exportJobs() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const vaultName = "NetSDKTestRsVault";
  const resourceGroupName = "SwaggerTestRg";
  const credential = new DefaultAzureCredential();
  const client = new RecoveryServicesBackupClient(credential, subscriptionId);
  const result = await client.jobs.export(vaultName, resourceGroupName);
  console.log(result);
}

exportJobs().catch(console.error);
