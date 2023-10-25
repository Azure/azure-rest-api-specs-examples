const { CdnManagementClient } = require("@azure/arm-cdn");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes an existing route with the specified route name under the specified subscription, resource group, profile, and AzureFrontDoor endpoint.
 *
 * @summary Deletes an existing route with the specified route name under the specified subscription, resource group, profile, and AzureFrontDoor endpoint.
 * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2023-05-01/examples/Routes_Delete.json
 */
async function routesDelete() {
  const subscriptionId = process.env["CDN_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["CDN_RESOURCE_GROUP"] || "RG";
  const profileName = "profile1";
  const endpointName = "endpoint1";
  const routeName = "route1";
  const credential = new DefaultAzureCredential();
  const client = new CdnManagementClient(credential, subscriptionId);
  const result = await client.routes.beginDeleteAndWait(
    resourceGroupName,
    profileName,
    endpointName,
    routeName
  );
  console.log(result);
}
