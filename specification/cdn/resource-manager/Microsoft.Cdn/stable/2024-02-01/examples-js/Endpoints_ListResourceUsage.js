const { CdnManagementClient } = require("@azure/arm-cdn");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Checks the quota and usage of geo filters and custom domains under the given endpoint.
 *
 * @summary Checks the quota and usage of geo filters and custom domains under the given endpoint.
 * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2024-02-01/examples/Endpoints_ListResourceUsage.json
 */
async function endpointsListResourceUsage() {
  const subscriptionId = process.env["CDN_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["CDN_RESOURCE_GROUP"] || "RG";
  const profileName = "profile1";
  const endpointName = "endpoint1";
  const credential = new DefaultAzureCredential();
  const client = new CdnManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.endpoints.listResourceUsage(
    resourceGroupName,
    profileName,
    endpointName,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
