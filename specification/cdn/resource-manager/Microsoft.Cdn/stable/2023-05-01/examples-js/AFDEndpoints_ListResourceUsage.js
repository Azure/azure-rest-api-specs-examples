const { CdnManagementClient } = require("@azure/arm-cdn");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Checks the quota and actual usage of endpoints under the given Azure Front Door profile.
 *
 * @summary Checks the quota and actual usage of endpoints under the given Azure Front Door profile.
 * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2023-05-01/examples/AFDEndpoints_ListResourceUsage.json
 */
async function afdEndpointsListResourceUsage() {
  const subscriptionId = process.env["CDN_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["CDN_RESOURCE_GROUP"] || "RG";
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
