const { ApplicationInsightsManagementClient } = require("@azure/arm-appinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets default work item configurations that exist for the application
 *
 * @summary Gets default work item configurations that exist for the application
 * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/WorkItemConfigDefaultGet.json
 */
async function workItemConfigurationsGetDefault() {
  const subscriptionId = "subid";
  const resourceGroupName = "my-resource-group";
  const resourceName = "my-component";
  const credential = new DefaultAzureCredential();
  const client = new ApplicationInsightsManagementClient(credential, subscriptionId);
  const result = await client.workItemConfigurations.getDefault(resourceGroupName, resourceName);
  console.log(result);
}

workItemConfigurationsGetDefault().catch(console.error);
