const { SiteRecoveryManagementClient } = require("@azure/arm-recoveryservices-siterecovery");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the data of supported operating systems by SRS.
 *
 * @summary Gets the data of supported operating systems by SRS.
 * x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2022-10-01/examples/SupportedOperatingSystems_Get.json
 */
async function getsTheDataOfSupportedOperatingSystemsBySrs() {
  const subscriptionId = "c183865e-6077-46f2-a3b1-deb0f4f4650a";
  const resourceName = "vault1";
  const resourceGroupName = "resourceGroupPS1";
  const credential = new DefaultAzureCredential();
  const client = new SiteRecoveryManagementClient(credential, subscriptionId);
  const result = await client.supportedOperatingSystemsOperations.get(
    resourceName,
    resourceGroupName
  );
  console.log(result);
}

getsTheDataOfSupportedOperatingSystemsBySrs().catch(console.error);
