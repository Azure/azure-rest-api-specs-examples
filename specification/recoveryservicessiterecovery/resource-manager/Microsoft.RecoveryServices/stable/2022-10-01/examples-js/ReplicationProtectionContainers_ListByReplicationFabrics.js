const { SiteRecoveryManagementClient } = require("@azure/arm-recoveryservices-siterecovery");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the protection containers in the specified fabric.
 *
 * @summary Lists the protection containers in the specified fabric.
 * x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2022-10-01/examples/ReplicationProtectionContainers_ListByReplicationFabrics.json
 */
async function getsTheListOfProtectionContainerForAFabric() {
  const subscriptionId = "c183865e-6077-46f2-a3b1-deb0f4f4650a";
  const resourceName = "vault1";
  const resourceGroupName = "resourceGroupPS1";
  const fabricName = "cloud1";
  const credential = new DefaultAzureCredential();
  const client = new SiteRecoveryManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.replicationProtectionContainers.listByReplicationFabrics(
    resourceName,
    resourceGroupName,
    fabricName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

getsTheListOfProtectionContainerForAFabric().catch(console.error);
