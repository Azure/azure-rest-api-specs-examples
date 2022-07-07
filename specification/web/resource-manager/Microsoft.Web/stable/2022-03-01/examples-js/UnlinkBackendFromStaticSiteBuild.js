const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Unlink a backend from a static site build
 *
 * @summary Unlink a backend from a static site build
 * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2022-03-01/examples/UnlinkBackendFromStaticSiteBuild.json
 */
async function unlinkABackendFromAStaticSiteBuild() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "rg";
  const name = "testStaticSite0";
  const environmentName = "12";
  const linkedBackendName = "testBackend";
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const result = await client.staticSites.unlinkBackendFromBuild(
    resourceGroupName,
    name,
    environmentName,
    linkedBackendName
  );
  console.log(result);
}

unlinkABackendFromAStaticSiteBuild().catch(console.error);
