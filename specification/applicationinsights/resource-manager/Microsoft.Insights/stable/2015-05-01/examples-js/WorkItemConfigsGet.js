const { ApplicationInsightsManagementClient } = require("@azure/arm-appinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the list work item configurations that exist for the application
 *
 * @summary Gets the list work item configurations that exist for the application
 * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/WorkItemConfigsGet.json
 */
async function workItemConfigurationsList() {
  const subscriptionId = "subid";
  const resourceGroupName = "my-resource-group";
  const resourceName = "my-component";
  const credential = new DefaultAzureCredential();
  const client = new ApplicationInsightsManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.workItemConfigurations.list(resourceGroupName, resourceName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

workItemConfigurationsList().catch(console.error);
