const { CdnManagementClient } = require("@azure/arm-cdn");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Check the name availability of a host name.
 *
 * @summary Check the name availability of a host name.
 * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/AFDProfiles_CheckHostNameAvailability.json
 */
async function afdProfilesCheckHostNameAvailability() {
  const subscriptionId = "subid";
  const resourceGroupName = "RG";
  const profileName = "profile1";
  const checkHostNameAvailabilityInput = {
    hostName: "www.someDomain.net",
  };
  const credential = new DefaultAzureCredential();
  const client = new CdnManagementClient(credential, subscriptionId);
  const result = await client.afdProfiles.checkHostNameAvailability(
    resourceGroupName,
    profileName,
    checkHostNameAvailabilityInput
  );
  console.log(result);
}

afdProfilesCheckHostNameAvailability().catch(console.error);
