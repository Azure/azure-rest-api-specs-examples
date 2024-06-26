const { ApplicationInsightsManagementClient } = require("@azure/arm-appinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete a work item configuration of an Application Insights component.
 *
 * @summary Delete a work item configuration of an Application Insights component.
 * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/WorkItemConfigDelete.json
 */
async function workItemConfigurationDelete() {
  const subscriptionId = "subid";
  const resourceGroupName = "my-resource-group";
  const resourceName = "my-component";
  const workItemConfigId = "Visual Studio Team Services";
  const credential = new DefaultAzureCredential();
  const client = new ApplicationInsightsManagementClient(credential, subscriptionId);
  const result = await client.workItemConfigurations.delete(
    resourceGroupName,
    resourceName,
    workItemConfigId
  );
  console.log(result);
}

workItemConfigurationDelete().catch(console.error);
