const { ApplicationInsightsManagementClient } = require("@azure/arm-appinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a single workbook by its resourceName.
 *
 * @summary Get a single workbook by its resourceName.
 * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2022-04-01/examples/WorkbookManagedGet.json
 */
async function workbookManagedGet() {
  const subscriptionId = "6b643656-33eb-422f-aee8-3ac145d124af";
  const resourceGroupName = "my-resource-group";
  const resourceName = "deadb33f-5e0d-4064-8ebb-1a4ed0313eb2";
  const credential = new DefaultAzureCredential();
  const client = new ApplicationInsightsManagementClient(credential, subscriptionId);
  const result = await client.workbooks.get(resourceGroupName, resourceName);
  console.log(result);
}

workbookManagedGet().catch(console.error);
