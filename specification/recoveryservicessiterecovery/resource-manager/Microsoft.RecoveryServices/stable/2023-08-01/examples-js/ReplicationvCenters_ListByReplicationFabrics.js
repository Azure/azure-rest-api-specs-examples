const { SiteRecoveryManagementClient } = require("@azure/arm-recoveryservices-siterecovery");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the vCenter servers registered in a fabric.
 *
 * @summary Lists the vCenter servers registered in a fabric.
 * x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-08-01/examples/ReplicationvCenters_ListByReplicationFabrics.json
 */
async function getsTheListOfVCenterRegisteredUnderAFabric() {
  const subscriptionId =
    process.env["RECOVERYSERVICESSITERECOVERY_SUBSCRIPTION_ID"] ||
    "7c943c1b-5122-4097-90c8-861411bdd574";
  const resourceName = "MadhaviVault";
  const resourceGroupName =
    process.env["RECOVERYSERVICESSITERECOVERY_RESOURCE_GROUP"] || "MadhaviVRG";
  const fabricName = "MadhaviFabric";
  const credential = new DefaultAzureCredential();
  const client = new SiteRecoveryManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.replicationvCenters.listByReplicationFabrics(
    resourceName,
    resourceGroupName,
    fabricName,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
