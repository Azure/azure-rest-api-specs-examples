const { RecoveryServicesClient } = require("@azure/arm-recoveryservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns a specified private link resource that need to be created for Backup and SiteRecovery
 *
 * @summary Returns a specified private link resource that need to be created for Backup and SiteRecovery
 * x-ms-original-file: specification/recoveryservices/resource-manager/Microsoft.RecoveryServices/stable/2022-04-01/examples/GetPrivateLinkResources.json
 */
async function getPrivateLinkResource() {
  const subscriptionId = "6c48fa17-39c7-45f1-90ac-47a587128ace";
  const resourceGroupName = "petesting";
  const vaultName = "pemsi-ecy-rsv2";
  const privateLinkResourceName = "backupResource";
  const credential = new DefaultAzureCredential();
  const client = new RecoveryServicesClient(credential, subscriptionId);
  const result = await client.privateLinkResourcesOperations.get(
    resourceGroupName,
    vaultName,
    privateLinkResourceName
  );
  console.log(result);
}

getPrivateLinkResource().catch(console.error);
