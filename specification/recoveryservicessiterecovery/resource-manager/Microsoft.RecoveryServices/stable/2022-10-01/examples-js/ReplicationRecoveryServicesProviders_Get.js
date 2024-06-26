const { SiteRecoveryManagementClient } = require("@azure/arm-recoveryservices-siterecovery");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the details of registered recovery services provider.
 *
 * @summary Gets the details of registered recovery services provider.
 * x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2022-10-01/examples/ReplicationRecoveryServicesProviders_Get.json
 */
async function getsTheDetailsOfARecoveryServicesProvider() {
  const subscriptionId = "c183865e-6077-46f2-a3b1-deb0f4f4650a";
  const resourceName = "vault1";
  const resourceGroupName = "resourceGroupPS1";
  const fabricName = "cloud1";
  const providerName = "241641e6-ee7b-4ee4-8141-821fadda43fa";
  const credential = new DefaultAzureCredential();
  const client = new SiteRecoveryManagementClient(credential, subscriptionId);
  const result = await client.replicationRecoveryServicesProviders.get(
    resourceName,
    resourceGroupName,
    fabricName,
    providerName
  );
  console.log(result);
}

getsTheDetailsOfARecoveryServicesProvider().catch(console.error);
