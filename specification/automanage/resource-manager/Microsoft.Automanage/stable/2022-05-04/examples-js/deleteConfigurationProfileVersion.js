const { AutomanageClient } = require("@azure/arm-automanage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete a configuration profile version
 *
 * @summary Delete a configuration profile version
 * x-ms-original-file: specification/automanage/resource-manager/Microsoft.Automanage/stable/2022-05-04/examples/deleteConfigurationProfileVersion.json
 */
async function deleteAConfigurationProfileVersion() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg";
  const configurationProfileName = "customConfigurationProfile";
  const versionName = "version1";
  const credential = new DefaultAzureCredential();
  const client = new AutomanageClient(credential, subscriptionId);
  const result = await client.configurationProfilesVersions.delete(
    resourceGroupName,
    configurationProfileName,
    versionName
  );
  console.log(result);
}

deleteAConfigurationProfileVersion().catch(console.error);
