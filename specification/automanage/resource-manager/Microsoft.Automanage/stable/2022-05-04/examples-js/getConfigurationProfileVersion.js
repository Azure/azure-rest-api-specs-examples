const { AutomanageClient } = require("@azure/arm-automanage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get information about a configuration profile version
 *
 * @summary Get information about a configuration profile version
 * x-ms-original-file: specification/automanage/resource-manager/Microsoft.Automanage/stable/2022-05-04/examples/getConfigurationProfileVersion.json
 */
async function getAConfigurationProfileVersion() {
  const subscriptionId = "mySubscriptionId";
  const configurationProfileName = "customConfigurationProfile";
  const versionName = "version1";
  const resourceGroupName = "myResourceGroupName";
  const credential = new DefaultAzureCredential();
  const client = new AutomanageClient(credential, subscriptionId);
  const result = await client.configurationProfilesVersions.get(
    configurationProfileName,
    versionName,
    resourceGroupName
  );
  console.log(result);
}

getAConfigurationProfileVersion().catch(console.error);
