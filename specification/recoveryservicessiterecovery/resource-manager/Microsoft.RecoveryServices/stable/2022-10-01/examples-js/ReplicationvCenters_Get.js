const { SiteRecoveryManagementClient } = require("@azure/arm-recoveryservices-siterecovery");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the details of a registered vCenter server(Add vCenter server).
 *
 * @summary Gets the details of a registered vCenter server(Add vCenter server).
 * x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2022-10-01/examples/ReplicationvCenters_Get.json
 */
async function getsTheDetailsOfAVCenter() {
  const subscriptionId = "7c943c1b-5122-4097-90c8-861411bdd574";
  const resourceName = "MadhaviVault";
  const resourceGroupName = "MadhaviVRG";
  const fabricName = "MadhaviFabric";
  const vcenterName = "esx-78";
  const credential = new DefaultAzureCredential();
  const client = new SiteRecoveryManagementClient(credential, subscriptionId);
  const result = await client.replicationvCenters.get(
    resourceName,
    resourceGroupName,
    fabricName,
    vcenterName
  );
  console.log(result);
}

getsTheDetailsOfAVCenter().catch(console.error);
