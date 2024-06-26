const { SiteRecoveryManagementClient } = require("@azure/arm-recoveryservices-siterecovery");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The operation to export the details of the Azure Site Recovery jobs of the vault.
 *
 * @summary The operation to export the details of the Azure Site Recovery jobs of the vault.
 * x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2022-10-01/examples/ReplicationJobs_Export.json
 */
async function exportsTheDetailsOfTheAzureSiteRecoveryJobsOfTheVault() {
  const subscriptionId = "c183865e-6077-46f2-a3b1-deb0f4f4650a";
  const resourceName = "vault1";
  const resourceGroupName = "resourceGroupPS1";
  const jobQueryParameter = {
    affectedObjectTypes: "",
    endTime: "2017-05-04T14:26:51.9161395Z",
    jobStatus: "",
    startTime: "2017-04-27T14:26:51.9161395Z",
  };
  const credential = new DefaultAzureCredential();
  const client = new SiteRecoveryManagementClient(credential, subscriptionId);
  const result = await client.replicationJobs.beginExportAndWait(
    resourceName,
    resourceGroupName,
    jobQueryParameter
  );
  console.log(result);
}

exportsTheDetailsOfTheAzureSiteRecoveryJobsOfTheVault().catch(console.error);
