const { CustomerInsightsManagementClient } = require("@azure/arm-customerinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a widget type in the specified hub.
 *
 * @summary Gets a widget type in the specified hub.
 * x-ms-original-file: specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/WidgetTypesGet.json
 */
async function widgetTypesGet() {
  const subscriptionId = "subid";
  const resourceGroupName = "TestHubRG";
  const hubName = "sdkTestHub";
  const widgetTypeName = "ActivityGauge";
  const credential = new DefaultAzureCredential();
  const client = new CustomerInsightsManagementClient(credential, subscriptionId);
  const result = await client.widgetTypes.get(resourceGroupName, hubName, widgetTypeName);
  console.log(result);
}

widgetTypesGet().catch(console.error);
