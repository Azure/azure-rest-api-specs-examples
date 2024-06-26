const { SiteRecoveryManagementClient } = require("@azure/arm-recoveryservices-siterecovery");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the available recovery points for a replication protected item.
 *
 * @summary Lists the available recovery points for a replication protected item.
 * x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2022-10-01/examples/RecoveryPoints_ListByReplicationProtectedItems.json
 */
async function getsTheListOfRecoveryPointsForAReplicationProtectedItem() {
  const subscriptionId = "c183865e-6077-46f2-a3b1-deb0f4f4650a";
  const resourceName = "vault1";
  const resourceGroupName = "resourceGroupPS1";
  const fabricName = "cloud1";
  const protectionContainerName = "cloud_6d224fc6-f326-5d35-96de-fbf51efb3179";
  const replicatedProtectedItemName = "f8491e4f-817a-40dd-a90c-af773978c75b";
  const credential = new DefaultAzureCredential();
  const client = new SiteRecoveryManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.recoveryPoints.listByReplicationProtectedItems(
    resourceName,
    resourceGroupName,
    fabricName,
    protectionContainerName,
    replicatedProtectedItemName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

getsTheListOfRecoveryPointsForAReplicationProtectedItem().catch(console.error);
