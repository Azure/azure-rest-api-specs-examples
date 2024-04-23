const { CdnManagementClient } = require("@azure/arm-cdn");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists existing AzureFrontDoor endpoints.
 *
 * @summary Lists existing AzureFrontDoor endpoints.
 * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2024-02-01/examples/AFDEndpoints_ListByProfile.json
 */
async function afdEndpointsListByProfile() {
  const subscriptionId = process.env["CDN_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["CDN_RESOURCE_GROUP"] || "RG";
  const profileName = "profile1";
  const credential = new DefaultAzureCredential();
  const client = new CdnManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.afdEndpoints.listByProfile(resourceGroupName, profileName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
