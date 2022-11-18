const { ApplicationInsightsManagementClient } = require("@azure/arm-appinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a list of Continuous Export configuration of an Application Insights component.
 *
 * @summary Gets a list of Continuous Export configuration of an Application Insights component.
 * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/ExportConfigurationsList.json
 */
async function exportConfigurationsList() {
  const subscriptionId = "subid";
  const resourceGroupName = "my-resource-group";
  const resourceName = "my-component";
  const credential = new DefaultAzureCredential();
  const client = new ApplicationInsightsManagementClient(credential, subscriptionId);
  const result = await client.exportConfigurations.list(resourceGroupName, resourceName);
  console.log(result);
}

exportConfigurationsList().catch(console.error);
