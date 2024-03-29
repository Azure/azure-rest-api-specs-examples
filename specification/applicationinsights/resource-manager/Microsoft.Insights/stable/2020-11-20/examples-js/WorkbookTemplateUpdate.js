const { ApplicationInsightsManagementClient } = require("@azure/arm-appinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates a workbook template that has already been added.
 *
 * @summary Updates a workbook template that has already been added.
 * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2020-11-20/examples/WorkbookTemplateUpdate.json
 */
async function workbookTemplateUpdate() {
  const subscriptionId = "6b643656-33eb-422f-aee8-3ac145d124af";
  const resourceGroupName = "my-resource-group";
  const resourceName = "my-template-resource";
  const credential = new DefaultAzureCredential();
  const client = new ApplicationInsightsManagementClient(credential, subscriptionId);
  const result = await client.workbookTemplates.update(resourceGroupName, resourceName);
  console.log(result);
}

workbookTemplateUpdate().catch(console.error);
