const { SiteRecoveryManagementClient } = require("@azure/arm-recoveryservices-siterecovery");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the list of ASR replication protected items in the protection container.
 *
 * @summary Gets the list of ASR replication protected items in the protection container.
 * x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-06-01/examples/ReplicationProtectedItems_ListByReplicationProtectionContainers.json
 */
async function getsTheListOfReplicationProtectedItems() {
  const subscriptionId =
    process.env["RECOVERYSERVICESSITERECOVERY_SUBSCRIPTION_ID"] ||
    "c183865e-6077-46f2-a3b1-deb0f4f4650a";
  const resourceName = "vault1";
  const resourceGroupName =
    process.env["RECOVERYSERVICESSITERECOVERY_RESOURCE_GROUP"] || "resourceGroupPS1";
  const fabricName = "cloud1";
  const protectionContainerName = "cloud_6d224fc6-f326-5d35-96de-fbf51efb3179";
  const credential = new DefaultAzureCredential();
  const client = new SiteRecoveryManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.replicationProtectedItems.listByReplicationProtectionContainers(
    resourceName,
    resourceGroupName,
    fabricName,
    protectionContainerName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
