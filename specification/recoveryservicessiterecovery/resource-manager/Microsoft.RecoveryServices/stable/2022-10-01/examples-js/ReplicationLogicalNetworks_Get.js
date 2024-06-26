const { SiteRecoveryManagementClient } = require("@azure/arm-recoveryservices-siterecovery");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the details of a logical network.
 *
 * @summary Gets the details of a logical network.
 * x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2022-10-01/examples/ReplicationLogicalNetworks_Get.json
 */
async function getsALogicalNetworkWithSpecifiedServerIdAndLogicalNetworkName() {
  const subscriptionId = "c183865e-6077-46f2-a3b1-deb0f4f4650a";
  const resourceName = "vault1";
  const resourceGroupName = "resourceGroupPS1";
  const fabricName = "cloud1";
  const logicalNetworkName = "87ab394f-165f-4aa9-bd84-b018500b4509";
  const credential = new DefaultAzureCredential();
  const client = new SiteRecoveryManagementClient(credential, subscriptionId);
  const result = await client.replicationLogicalNetworks.get(
    resourceName,
    resourceGroupName,
    fabricName,
    logicalNetworkName
  );
  console.log(result);
}

getsALogicalNetworkWithSpecifiedServerIdAndLogicalNetworkName().catch(console.error);
