const { ApplicationInsightsManagementClient } = require("@azure/arm-appinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Remove a favorite that is associated to an Application Insights component.
 *
 * @summary Remove a favorite that is associated to an Application Insights component.
 * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/FavoriteDelete.json
 */
async function favoriteList() {
  const subscriptionId = "subid";
  const resourceGroupName = "my-resource-group";
  const resourceName = "my-ai-component";
  const favoriteId = "deadb33f-5e0d-4064-8ebb-1a4ed0313eb2";
  const credential = new DefaultAzureCredential();
  const client = new ApplicationInsightsManagementClient(credential, subscriptionId);
  const result = await client.favorites.delete(resourceGroupName, resourceName, favoriteId);
  console.log(result);
}

favoriteList().catch(console.error);
