const { AutomanageClient } = require("@azure/arm-automanage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get information about a configuration profile
 *
 * @summary Get information about a configuration profile
 * x-ms-original-file: specification/automanage/resource-manager/Microsoft.Automanage/stable/2022-05-04/examples/getConfigurationProfile.json
 */
async function getAConfigurationProfile() {
  const subscriptionId = "mySubscriptionId";
  const configurationProfileName = "customConfigurationProfile";
  const resourceGroupName = "myResourceGroupName";
  const credential = new DefaultAzureCredential();
  const client = new AutomanageClient(credential, subscriptionId);
  const result = await client.configurationProfiles.get(
    configurationProfileName,
    resourceGroupName
  );
  console.log(result);
}

getAConfigurationProfile().catch(console.error);
