const { ApplicationInsightsManagementClient } = require("@azure/arm-appinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates a workbook that has already been added.
 *
 * @summary Updates a workbook that has already been added.
 * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2022-04-01/examples/WorkbookManagedUpdate.json
 */
async function workbookManagedUpdate() {
  const subscriptionId = "6b643656-33eb-422f-aee8-3ac145d124af";
  const resourceGroupName = "my-resource-group";
  const resourceName = "deadb33f-5e0d-4064-8ebb-1a4ed0313eb2";
  const sourceId =
    "/subscriptions/6b643656-33eb-422f-aee8-3ac145d124af/resourcegroups/my-resource-group";
  const options = { sourceId };
  const credential = new DefaultAzureCredential();
  const client = new ApplicationInsightsManagementClient(credential, subscriptionId);
  const result = await client.workbooks.update(resourceGroupName, resourceName, options);
  console.log(result);
}

workbookManagedUpdate().catch(console.error);
