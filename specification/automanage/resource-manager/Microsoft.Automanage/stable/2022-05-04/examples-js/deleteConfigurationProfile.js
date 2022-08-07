const { AutomanageClient } = require("@azure/arm-automanage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete a configuration profile
 *
 * @summary Delete a configuration profile
 * x-ms-original-file: specification/automanage/resource-manager/Microsoft.Automanage/stable/2022-05-04/examples/deleteConfigurationProfile.json
 */
async function deleteAConfigurationProfile() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg";
  const configurationProfileName = "customConfigurationProfile";
  const credential = new DefaultAzureCredential();
  const client = new AutomanageClient(credential, subscriptionId);
  const result = await client.configurationProfiles.delete(
    resourceGroupName,
    configurationProfileName
  );
  console.log(result);
}

deleteAConfigurationProfile().catch(console.error);
