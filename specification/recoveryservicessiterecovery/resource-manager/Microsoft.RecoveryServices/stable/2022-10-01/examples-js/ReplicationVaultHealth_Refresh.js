const { SiteRecoveryManagementClient } = require("@azure/arm-recoveryservices-siterecovery");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Refreshes health summary of the vault.
 *
 * @summary Refreshes health summary of the vault.
 * x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2022-10-01/examples/ReplicationVaultHealth_Refresh.json
 */
async function refreshesHealthSummaryOfTheVault() {
  const subscriptionId = "c183865e-6077-46f2-a3b1-deb0f4f4650a";
  const resourceName = "vault1";
  const resourceGroupName = "resourceGroupPS1";
  const credential = new DefaultAzureCredential();
  const client = new SiteRecoveryManagementClient(credential, subscriptionId);
  const result = await client.replicationVaultHealth.beginRefreshAndWait(
    resourceName,
    resourceGroupName
  );
  console.log(result);
}

refreshesHealthSummaryOfTheVault().catch(console.error);
