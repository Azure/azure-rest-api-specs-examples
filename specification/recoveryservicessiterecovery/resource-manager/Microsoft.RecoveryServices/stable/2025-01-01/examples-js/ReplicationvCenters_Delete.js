const { SiteRecoveryManagementClient } = require("@azure/arm-recoveryservices-siterecovery");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to The operation to remove(unregister) a registered vCenter server from the vault.
 *
 * @summary The operation to remove(unregister) a registered vCenter server from the vault.
 * x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples/ReplicationvCenters_Delete.json
 */
async function removeVCenterOperation() {
  const subscriptionId =
    process.env["RECOVERYSERVICESSITERECOVERY_SUBSCRIPTION_ID"] ||
    "7c943c1b-5122-4097-90c8-861411bdd574";
  const resourceGroupName =
    process.env["RECOVERYSERVICESSITERECOVERY_RESOURCE_GROUP"] || "MadhaviVRG";
  const resourceName = "MadhaviVault";
  const fabricName = "MadhaviFabric";
  const vcenterName = "esx-78";
  const credential = new DefaultAzureCredential();
  const client = new SiteRecoveryManagementClient(credential, subscriptionId);
  const result = await client.replicationvCenters.beginDeleteAndWait(
    resourceGroupName,
    resourceName,
    fabricName,
    vcenterName,
  );
  console.log(result);
}
