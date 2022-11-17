const { ApplicationInsightsManagementClient } = require("@azure/arm-appinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a single workbook template by its resourceName.
 *
 * @summary Get a single workbook template by its resourceName.
 * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2020-11-20/examples/WorkbookTemplateGet.json
 */
async function workbookTemplateGet() {
  const subscriptionId = "6b643656-33eb-422f-aee8-3ac145d124af";
  const resourceGroupName = "my-resource-group";
  const resourceName = "my-resource-name";
  const credential = new DefaultAzureCredential();
  const client = new ApplicationInsightsManagementClient(credential, subscriptionId);
  const result = await client.workbookTemplates.get(resourceGroupName, resourceName);
  console.log(result);
}

workbookTemplateGet().catch(console.error);
