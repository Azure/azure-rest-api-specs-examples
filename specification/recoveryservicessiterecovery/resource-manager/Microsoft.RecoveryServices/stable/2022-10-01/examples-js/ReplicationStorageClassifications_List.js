const { SiteRecoveryManagementClient } = require("@azure/arm-recoveryservices-siterecovery");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the storage classifications in the vault.
 *
 * @summary Lists the storage classifications in the vault.
 * x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2022-10-01/examples/ReplicationStorageClassifications_List.json
 */
async function getsTheListOfStorageClassificationObjectsUnderAVault() {
  const subscriptionId = "9112a37f-0f3e-46ec-9c00-060c6edca071";
  const resourceName = "vault1";
  const resourceGroupName = "resourceGroupPS1";
  const credential = new DefaultAzureCredential();
  const client = new SiteRecoveryManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.replicationStorageClassifications.list(
    resourceName,
    resourceGroupName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

getsTheListOfStorageClassificationObjectsUnderAVault().catch(console.error);
