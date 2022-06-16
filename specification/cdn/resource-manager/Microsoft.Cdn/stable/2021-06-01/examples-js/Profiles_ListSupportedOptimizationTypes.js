const { CdnManagementClient } = require("@azure/arm-cdn");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the supported optimization types for the current profile. A user can create an endpoint with an optimization type from the listed values.
 *
 * @summary Gets the supported optimization types for the current profile. A user can create an endpoint with an optimization type from the listed values.
 * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/Profiles_ListSupportedOptimizationTypes.json
 */
async function profilesListSupportedOptimizationTypes() {
  const subscriptionId = "subid";
  const resourceGroupName = "RG";
  const profileName = "profile1";
  const credential = new DefaultAzureCredential();
  const client = new CdnManagementClient(credential, subscriptionId);
  const result = await client.profiles.listSupportedOptimizationTypes(
    resourceGroupName,
    profileName
  );
  console.log(result);
}

profilesListSupportedOptimizationTypes().catch(console.error);
