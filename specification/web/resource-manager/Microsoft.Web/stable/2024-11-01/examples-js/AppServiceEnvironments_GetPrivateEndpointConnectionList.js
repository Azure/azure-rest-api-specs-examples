const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Description for Gets the list of private endpoints associated with a hosting environment
 *
 * @summary Description for Gets the list of private endpoints associated with a hosting environment
 * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2024-11-01/examples/AppServiceEnvironments_GetPrivateEndpointConnectionList.json
 */
async function getsTheListOfPrivateEndpointsAssociatedWithAHostingEnvironment() {
  const subscriptionId =
    process.env["APPSERVICE_SUBSCRIPTION_ID"] || "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = process.env["APPSERVICE_RESOURCE_GROUP"] || "test-rg";
  const name = "test-ase";
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.appServiceEnvironments.listPrivateEndpointConnectionList(
    resourceGroupName,
    name,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
