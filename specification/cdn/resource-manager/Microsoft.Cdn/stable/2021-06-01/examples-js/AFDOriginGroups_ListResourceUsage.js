const { CdnManagementClient } = require("@azure/arm-cdn");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Checks the quota and actual usage of the given AzureFrontDoor origin group under the given CDN profile.
 *
 * @summary Checks the quota and actual usage of the given AzureFrontDoor origin group under the given CDN profile.
 * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/AFDOriginGroups_ListResourceUsage.json
 */
async function afdOriginGroupsListResourceUsage() {
  const subscriptionId = "subid";
  const resourceGroupName = "RG";
  const profileName = "profile1";
  const originGroupName = "origingroup1";
  const credential = new DefaultAzureCredential();
  const client = new CdnManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.afdOriginGroups.listResourceUsage(
    resourceGroupName,
    profileName,
    originGroupName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

afdOriginGroupsListResourceUsage().catch(console.error);
