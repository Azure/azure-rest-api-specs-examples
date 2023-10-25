const { CdnManagementClient } = require("@azure/arm-cdn");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all of the existing origins within a profile.
 *
 * @summary Lists all of the existing origins within a profile.
 * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2023-05-01/examples/Routes_ListByEndpoint.json
 */
async function routesListByEndpoint() {
  const subscriptionId = process.env["CDN_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["CDN_RESOURCE_GROUP"] || "RG";
  const profileName = "profile1";
  const endpointName = "endpoint1";
  const credential = new DefaultAzureCredential();
  const client = new CdnManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.routes.listByEndpoint(
    resourceGroupName,
    profileName,
    endpointName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
