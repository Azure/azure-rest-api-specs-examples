const { ApplicationInsightsManagementClient } = require("@azure/arm-appinsights");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Get the ProactiveDetection configuration for this configuration id.
 *
 * @summary Get the ProactiveDetection configuration for this configuration id.
 * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/ProactiveDetectionConfigurationGet.json
 */
async function proactiveDetectionConfigurationGet() {
  const subscriptionId = process.env["APPLICATIONINSIGHTS_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName =
    process.env["APPLICATIONINSIGHTS_RESOURCE_GROUP"] || "my-resource-group";
  const resourceName = "my-component";
  const configurationId = "slowpageloadtime";
  const credential = new DefaultAzureCredential();
  const client = new ApplicationInsightsManagementClient(credential, subscriptionId);
  const result = await client.proactiveDetectionConfigurations.get(
    resourceGroupName,
    resourceName,
    configurationId,
  );
  console.log(result);
}
