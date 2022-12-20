const { SiteRecoveryManagementClient } = require("@azure/arm-recoveryservices-siterecovery");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the list of vault setting. This includes the Migration Hub connection settings.
 *
 * @summary Gets the list of vault setting. This includes the Migration Hub connection settings.
 * x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2022-10-01/examples/ReplicationVaultSetting_List.json
 */
async function getsTheListOfVaultSetting() {
  const subscriptionId = "c183865e-6077-46f2-a3b1-deb0f4f4650a";
  const resourceName = "vault1";
  const resourceGroupName = "resourceGroupPS1";
  const credential = new DefaultAzureCredential();
  const client = new SiteRecoveryManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.replicationVaultSetting.list(resourceName, resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

getsTheListOfVaultSetting().catch(console.error);
