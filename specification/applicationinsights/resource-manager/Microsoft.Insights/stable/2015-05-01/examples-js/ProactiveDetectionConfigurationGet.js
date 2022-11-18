const { ApplicationInsightsManagementClient } = require("@azure/arm-appinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the ProactiveDetection configuration for this configuration id.
 *
 * @summary Get the ProactiveDetection configuration for this configuration id.
 * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/ProactiveDetectionConfigurationGet.json
 */
async function proactiveDetectionConfigurationGet() {
  const subscriptionId = "subid";
  const resourceGroupName = "my-resource-group";
  const resourceName = "my-component";
  const configurationId = "slowpageloadtime";
  const credential = new DefaultAzureCredential();
  const client = new ApplicationInsightsManagementClient(credential, subscriptionId);
  const result = await client.proactiveDetectionConfigurations.get(
    resourceGroupName,
    resourceName,
    configurationId
  );
  console.log(result);
}

proactiveDetectionConfigurationGet().catch(console.error);
