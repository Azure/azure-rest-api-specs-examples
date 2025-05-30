const { ApplicationInsightsManagementClient } = require("@azure/arm-appinsights");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Create a work item configuration for an Application Insights component.
 *
 * @summary Create a work item configuration for an Application Insights component.
 * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/WorkItemConfigCreate.json
 */
async function workItemConfigurationsCreate() {
  const subscriptionId = process.env["APPLICATIONINSIGHTS_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName =
    process.env["APPLICATIONINSIGHTS_RESOURCE_GROUP"] || "my-resource-group";
  const resourceName = "my-component";
  const workItemConfigurationProperties = {
    connectorDataConfiguration:
      '{"VSOAccountBaseUrl":"https://testtodelete.visualstudio.com","ProjectCollection":"DefaultCollection","Project":"todeletefirst","ResourceId":"d0662b05-439a-4a1b-840b-33a7f8b42ebf","Custom":"{\\"/fields/System.WorkItemType\\":\\"Bug\\",\\"/fields/System.AreaPath\\":\\"todeletefirst\\",\\"/fields/System.AssignedTo\\":\\"\\"}"}',
    connectorId: "d334e2a4-6733-488e-8645-a9fdc1694f41",
    validateOnly: true,
    workItemProperties: { 0: "[object Object]", 1: "[object Object]" },
  };
  const credential = new DefaultAzureCredential();
  const client = new ApplicationInsightsManagementClient(credential, subscriptionId);
  const result = await client.workItemConfigurations.create(
    resourceGroupName,
    resourceName,
    workItemConfigurationProperties,
  );
  console.log(result);
}
