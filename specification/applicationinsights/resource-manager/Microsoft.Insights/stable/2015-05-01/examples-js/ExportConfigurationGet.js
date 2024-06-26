const { ApplicationInsightsManagementClient } = require("@azure/arm-appinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the Continuous Export configuration for this export id.
 *
 * @summary Get the Continuous Export configuration for this export id.
 * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/ExportConfigurationGet.json
 */
async function exportConfigurationGet() {
  const subscriptionId = "subid";
  const resourceGroupName = "my-resource-group";
  const resourceName = "my-component";
  const exportId = "uGOoki0jQsyEs3IdQ83Q4QsNr4=";
  const credential = new DefaultAzureCredential();
  const client = new ApplicationInsightsManagementClient(credential, subscriptionId);
  const result = await client.exportConfigurations.get(resourceGroupName, resourceName, exportId);
  console.log(result);
}

exportConfigurationGet().catch(console.error);
