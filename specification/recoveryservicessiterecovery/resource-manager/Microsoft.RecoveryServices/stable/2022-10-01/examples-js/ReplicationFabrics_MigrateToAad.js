const { SiteRecoveryManagementClient } = require("@azure/arm-recoveryservices-siterecovery");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The operation to migrate an Azure Site Recovery fabric to AAD.
 *
 * @summary The operation to migrate an Azure Site Recovery fabric to AAD.
 * x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2022-10-01/examples/ReplicationFabrics_MigrateToAad.json
 */
async function migratesTheSiteToAad() {
  const subscriptionId = "c183865e-6077-46f2-a3b1-deb0f4f4650a";
  const resourceName = "vault1";
  const resourceGroupName = "resourceGroupPS1";
  const fabricName = "cloud1";
  const credential = new DefaultAzureCredential();
  const client = new SiteRecoveryManagementClient(credential, subscriptionId);
  const result = await client.replicationFabrics.beginMigrateToAadAndWait(
    resourceName,
    resourceGroupName,
    fabricName
  );
  console.log(result);
}

migratesTheSiteToAad().catch(console.error);
