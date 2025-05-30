const { SiteRecoveryManagementClient } = require("@azure/arm-recoveryservices-siterecovery");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Gets the data of supported operating systems by SRS.
 *
 * @summary Gets the data of supported operating systems by SRS.
 * x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples/SupportedOperatingSystems_Get.json
 */
async function getsTheDataOfSupportedOperatingSystemsBySrs() {
  const subscriptionId =
    process.env["RECOVERYSERVICESSITERECOVERY_SUBSCRIPTION_ID"] ||
    "c183865e-6077-46f2-a3b1-deb0f4f4650a";
  const resourceGroupName =
    process.env["RECOVERYSERVICESSITERECOVERY_RESOURCE_GROUP"] || "resourceGroupPS1";
  const resourceName = "vault1";
  const credential = new DefaultAzureCredential();
  const client = new SiteRecoveryManagementClient(credential, subscriptionId);
  const result = await client.supportedOperatingSystemsOperations.get(
    resourceGroupName,
    resourceName,
  );
  console.log(result);
}
