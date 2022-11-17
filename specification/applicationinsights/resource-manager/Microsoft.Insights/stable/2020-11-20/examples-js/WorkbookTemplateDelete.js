const { ApplicationInsightsManagementClient } = require("@azure/arm-appinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete a workbook template.
 *
 * @summary Delete a workbook template.
 * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2020-11-20/examples/WorkbookTemplateDelete.json
 */
async function workbookTemplateDelete() {
  const subscriptionId = "subid";
  const resourceGroupName = "my-resource-group";
  const resourceName = "my-template-resource";
  const credential = new DefaultAzureCredential();
  const client = new ApplicationInsightsManagementClient(credential, subscriptionId);
  const result = await client.workbookTemplates.delete(resourceGroupName, resourceName);
  console.log(result);
}

workbookTemplateDelete().catch(console.error);
