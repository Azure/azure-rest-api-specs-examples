const { ApplicationInsightsManagementClient } = require("@azure/arm-appinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns feature capabilities of the application insights component.
 *
 * @summary Returns feature capabilities of the application insights component.
 * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/FeatureCapabilitiesGet.json
 */
async function componentCurrentBillingFeaturesGet() {
  const subscriptionId = "subid";
  const resourceGroupName = "my-resource-group";
  const resourceName = "my-component";
  const credential = new DefaultAzureCredential();
  const client = new ApplicationInsightsManagementClient(credential, subscriptionId);
  const result = await client.componentFeatureCapabilities.get(resourceGroupName, resourceName);
  console.log(result);
}

componentCurrentBillingFeaturesGet().catch(console.error);
