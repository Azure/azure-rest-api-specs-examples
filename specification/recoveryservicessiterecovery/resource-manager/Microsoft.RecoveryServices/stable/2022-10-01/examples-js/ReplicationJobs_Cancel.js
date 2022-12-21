const { SiteRecoveryManagementClient } = require("@azure/arm-recoveryservices-siterecovery");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The operation to cancel an Azure Site Recovery job.
 *
 * @summary The operation to cancel an Azure Site Recovery job.
 * x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2022-10-01/examples/ReplicationJobs_Cancel.json
 */
async function cancelsTheSpecifiedJob() {
  const subscriptionId = "c183865e-6077-46f2-a3b1-deb0f4f4650a";
  const resourceName = "vault1";
  const resourceGroupName = "resourceGroupPS1";
  const jobName = "2653c648-fc72-4316-86f3-fdf8eaa0066b";
  const credential = new DefaultAzureCredential();
  const client = new SiteRecoveryManagementClient(credential, subscriptionId);
  const result = await client.replicationJobs.beginCancelAndWait(
    resourceName,
    resourceGroupName,
    jobName
  );
  console.log(result);
}

cancelsTheSpecifiedJob().catch(console.error);
