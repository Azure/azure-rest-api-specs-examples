const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Description for Gets a private endpoint connection
 *
 * @summary Description for Gets a private endpoint connection
 * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2022-03-01/examples/GetSitePrivateEndpointConnection.json
 */
async function getAPrivateEndpointConnectionForASite() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "rg";
  const name = "testSite";
  const privateEndpointConnectionName = "connection";
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const result = await client.staticSites.getPrivateEndpointConnection(
    resourceGroupName,
    name,
    privateEndpointConnectionName
  );
  console.log(result);
}

getAPrivateEndpointConnectionForASite().catch(console.error);
