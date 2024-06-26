const { ApplicationInsightsManagementClient } = require("@azure/arm-appinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a list of Analytics Items defined within an Application Insights component.
 *
 * @summary Gets a list of Analytics Items defined within an Application Insights component.
 * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/AnalyticsItemList.json
 */
async function analyticsItemList() {
  const subscriptionId = "subid";
  const resourceGroupName = "my-resource-group";
  const resourceName = "my-component";
  const scopePath = "analyticsItems";
  const credential = new DefaultAzureCredential();
  const client = new ApplicationInsightsManagementClient(credential, subscriptionId);
  const result = await client.analyticsItems.list(resourceGroupName, resourceName, scopePath);
  console.log(result);
}

analyticsItemList().catch(console.error);
