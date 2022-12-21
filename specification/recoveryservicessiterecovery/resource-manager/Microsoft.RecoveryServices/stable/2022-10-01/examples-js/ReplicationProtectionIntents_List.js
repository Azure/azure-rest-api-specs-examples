const { SiteRecoveryManagementClient } = require("@azure/arm-recoveryservices-siterecovery");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the list of ASR replication protection intent objects in the vault.
 *
 * @summary Gets the list of ASR replication protection intent objects in the vault.
 * x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2022-10-01/examples/ReplicationProtectionIntents_List.json
 */
async function getsTheListOfReplicationProtectionIntentObjects() {
  const subscriptionId = "509099b2-9d2c-4636-b43e-bd5cafb6be69";
  const resourceName = "2007vttp";
  const resourceGroupName = "resourceGroupPS1";
  const credential = new DefaultAzureCredential();
  const client = new SiteRecoveryManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.replicationProtectionIntents.list(
    resourceName,
    resourceGroupName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

getsTheListOfReplicationProtectionIntentObjects().catch(console.error);
