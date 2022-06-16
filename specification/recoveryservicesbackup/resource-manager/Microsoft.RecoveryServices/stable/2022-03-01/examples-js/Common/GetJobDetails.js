const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets extended information associated with the job.
 *
 * @summary Gets extended information associated with the job.
 * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2022-03-01/examples/Common/GetJobDetails.json
 */
async function getJobDetails() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const vaultName = "NetSDKTestRsVault";
  const resourceGroupName = "SwaggerTestRg";
  const jobName = "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new RecoveryServicesBackupClient(credential, subscriptionId);
  const result = await client.jobDetails.get(vaultName, resourceGroupName, jobName);
  console.log(result);
}

getJobDetails().catch(console.error);
