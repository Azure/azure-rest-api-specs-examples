const { SiteRecoveryManagementClient } = require("@azure/arm-recoveryservices-siterecovery");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The operation to purge(force delete) a protection container mapping.
 *
 * @summary The operation to purge(force delete) a protection container mapping.
 * x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2022-10-01/examples/ReplicationProtectionContainerMappings_Purge.json
 */
async function purgeProtectionContainerMapping() {
  const subscriptionId = "c183865e-6077-46f2-a3b1-deb0f4f4650a";
  const resourceName = "vault1";
  const resourceGroupName = "resourceGroupPS1";
  const fabricName = "cloud1";
  const protectionContainerName = "cloud_6d224fc6-f326-5d35-96de-fbf51efb3179";
  const mappingName = "cloud1protectionprofile1";
  const credential = new DefaultAzureCredential();
  const client = new SiteRecoveryManagementClient(credential, subscriptionId);
  const result = await client.replicationProtectionContainerMappings.beginPurgeAndWait(
    resourceName,
    resourceGroupName,
    fabricName,
    protectionContainerName,
    mappingName
  );
  console.log(result);
}

purgeProtectionContainerMapping().catch(console.error);
