const { SiteRecoveryManagementClient } = require("@azure/arm-recoveryservices-siterecovery");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Gets the details of a protection container.
 *
 * @summary Gets the details of a protection container.
 * x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples/ReplicationProtectionContainers_Get.json
 */
async function getsTheProtectionContainerDetails() {
  const subscriptionId =
    process.env["RECOVERYSERVICESSITERECOVERY_SUBSCRIPTION_ID"] ||
    "c183865e-6077-46f2-a3b1-deb0f4f4650a";
  const resourceGroupName =
    process.env["RECOVERYSERVICESSITERECOVERY_RESOURCE_GROUP"] || "resourceGroupPS1";
  const resourceName = "vault1";
  const fabricName = "cloud1";
  const protectionContainerName = "cloud_6d224fc6-f326-5d35-96de-fbf51efb3179";
  const credential = new DefaultAzureCredential();
  const client = new SiteRecoveryManagementClient(credential, subscriptionId);
  const result = await client.replicationProtectionContainers.get(
    resourceGroupName,
    resourceName,
    fabricName,
    protectionContainerName,
  );
  console.log(result);
}
