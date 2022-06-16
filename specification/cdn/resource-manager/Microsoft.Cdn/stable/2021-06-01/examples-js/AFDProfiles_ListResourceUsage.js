const { CdnManagementClient } = require("@azure/arm-cdn");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Checks the quota and actual usage of endpoints under the given CDN profile.
 *
 * @summary Checks the quota and actual usage of endpoints under the given CDN profile.
 * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/AFDProfiles_ListResourceUsage.json
 */
async function afdProfilesListResourceUsage() {
  const subscriptionId = "subid";
  const resourceGroupName = "RG";
  const profileName = "profile1";
  const credential = new DefaultAzureCredential();
  const client = new CdnManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.afdProfiles.listResourceUsage(resourceGroupName, profileName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

afdProfilesListResourceUsage().catch(console.error);
