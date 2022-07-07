const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Description for Gets the list of private endpoint connections associated with a static site
 *
 * @summary Description for Gets the list of private endpoint connections associated with a static site
 * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2022-03-01/examples/GetSitePrivateEndpointConnectionList.json
 */
async function getAListOfPrivateEndpointConnectionsAssociatedWithASite() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "rg";
  const name = "testStaticSite0";
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.staticSites.listPrivateEndpointConnectionList(
    resourceGroupName,
    name
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

getAListOfPrivateEndpointConnectionsAssociatedWithASite().catch(console.error);
