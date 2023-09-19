const { SiteRecoveryManagementClient } = require("@azure/arm-recoveryservices-siterecovery");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all the logical networks of the Azure Site Recovery fabric.
 *
 * @summary Lists all the logical networks of the Azure Site Recovery fabric.
 * x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-06-01/examples/ReplicationLogicalNetworks_ListByReplicationFabrics.json
 */
async function getsTheListOfLogicalNetworksUnderAFabric() {
  const subscriptionId =
    process.env["RECOVERYSERVICESSITERECOVERY_SUBSCRIPTION_ID"] ||
    "c183865e-6077-46f2-a3b1-deb0f4f4650a";
  const resourceName = "vault1";
  const resourceGroupName =
    process.env["RECOVERYSERVICESSITERECOVERY_RESOURCE_GROUP"] || "resourceGroupPS1";
  const fabricName = "cloud1";
  const credential = new DefaultAzureCredential();
  const client = new SiteRecoveryManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.replicationLogicalNetworks.listByReplicationFabrics(
    resourceName,
    resourceGroupName,
    fabricName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
