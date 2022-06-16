const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Description for Register a user provided function app with a static site build
 *
 * @summary Description for Register a user provided function app with a static site build
 * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/RegisterUserProvidedFunctionAppWithStaticSiteBuild.json
 */
async function registerAUserProvidedFunctionAppWithAStaticSiteBuild() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "rg";
  const name = "testStaticSite0";
  const environmentName = "default";
  const functionAppName = "testFunctionApp";
  const isForced = true;
  const staticSiteUserProvidedFunctionEnvelope = {
    functionAppRegion: "West US 2",
    functionAppResourceId:
      "/subscription/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/functionRG/providers/Microsoft.Web/sites/testFunctionApp",
  };
  const options = {
    isForced,
  };
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const result =
    await client.staticSites.beginRegisterUserProvidedFunctionAppWithStaticSiteBuildAndWait(
      resourceGroupName,
      name,
      environmentName,
      functionAppName,
      staticSiteUserProvidedFunctionEnvelope,
      options
    );
  console.log(result);
}

registerAUserProvidedFunctionAppWithAStaticSiteBuild().catch(console.error);
