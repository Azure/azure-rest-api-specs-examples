const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns the details of a linked backend linked to a static site by name
 *
 * @summary Returns the details of a linked backend linked to a static site by name
 * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2022-03-01/examples/GetLinkedBackendForStaticSite.json
 */
async function getDetailsOfTheLinkedBackendRegisteredWithAStaticSiteByName() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "rg";
  const name = "testStaticSite0";
  const linkedBackendName = "testBackend";
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const result = await client.staticSites.getLinkedBackend(
    resourceGroupName,
    name,
    linkedBackendName
  );
  console.log(result);
}

getDetailsOfTheLinkedBackendRegisteredWithAStaticSiteByName().catch(console.error);
