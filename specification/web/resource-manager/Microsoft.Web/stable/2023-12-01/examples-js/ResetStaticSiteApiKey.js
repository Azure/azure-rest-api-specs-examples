const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Description for Resets the api key for an existing static site.
 *
 * @summary Description for Resets the api key for an existing static site.
 * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2023-12-01/examples/ResetStaticSiteApiKey.json
 */
async function resetTheApiKeyForAStaticSite() {
  const subscriptionId =
    process.env["APPSERVICE_SUBSCRIPTION_ID"] || "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = process.env["APPSERVICE_RESOURCE_GROUP"] || "rg";
  const name = "testStaticSite0";
  const resetPropertiesEnvelope = {
    repositoryToken: "repoToken123",
    shouldUpdateRepository: true,
  };
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const result = await client.staticSites.resetStaticSiteApiKey(
    resourceGroupName,
    name,
    resetPropertiesEnvelope,
  );
  console.log(result);
}
