const { ApplicationInsightsManagementClient } = require("@azure/arm-appinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete a private workbook.
 *
 * @summary Delete a private workbook.
 * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2021-03-08/examples/MyWorkbookDelete.json
 */
async function workbookDelete() {
  const subscriptionId = "subid";
  const resourceGroupName = "my-resource-group";
  const resourceName = "deadb33f-5e0d-4064-8ebb-1a4ed0313eb2";
  const credential = new DefaultAzureCredential();
  const client = new ApplicationInsightsManagementClient(credential, subscriptionId);
  const result = await client.myWorkbooks.delete(resourceGroupName, resourceName);
  console.log(result);
}

workbookDelete().catch(console.error);
