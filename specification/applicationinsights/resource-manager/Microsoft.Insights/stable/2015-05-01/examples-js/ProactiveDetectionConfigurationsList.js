const { ApplicationInsightsManagementClient } = require("@azure/arm-appinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a list of ProactiveDetection configurations of an Application Insights component.
 *
 * @summary Gets a list of ProactiveDetection configurations of an Application Insights component.
 * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/ProactiveDetectionConfigurationsList.json
 */
async function proactiveDetectionConfigurationsList() {
  const subscriptionId = "subid";
  const resourceGroupName = "my-resource-group";
  const resourceName = "my-component";
  const credential = new DefaultAzureCredential();
  const client = new ApplicationInsightsManagementClient(credential, subscriptionId);
  const result = await client.proactiveDetectionConfigurations.list(
    resourceGroupName,
    resourceName
  );
  console.log(result);
}

proactiveDetectionConfigurationsList().catch(console.error);
