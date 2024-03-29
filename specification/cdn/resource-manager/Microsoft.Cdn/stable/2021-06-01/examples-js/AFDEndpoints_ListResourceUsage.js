const { CdnManagementClient } = require("@azure/arm-cdn");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Checks the quota and actual usage of the given AzureFrontDoor endpoint under the given CDN profile.
 *
 * @summary Checks the quota and actual usage of the given AzureFrontDoor endpoint under the given CDN profile.
 * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/AFDEndpoints_ListResourceUsage.json
 */
async function afdEndpointsListResourceUsage() {
  const subscriptionId = "subid";
  const resourceGroupName = "RG";
  const profileName = "profile1";
  const endpointName = "endpoint1";
  const credential = new DefaultAzureCredential();
  const client = new CdnManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.afdEndpoints.listResourceUsage(
    resourceGroupName,
    profileName,
    endpointName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

afdEndpointsListResourceUsage().catch(console.error);
