const { ApplicationInsightsManagementClient } = require("@azure/arm-appinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get all private workbooks defined within a specified subscription and category.
 *
 * @summary Get all private workbooks defined within a specified subscription and category.
 * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2021-03-08/examples/MyWorkbooksList.json
 */
async function workbooksList() {
  const subscriptionId = "6b643656-33eb-422f-aee8-3ac145d124af";
  const category = "workbook";
  const credential = new DefaultAzureCredential();
  const client = new ApplicationInsightsManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.myWorkbooks.listBySubscription(category)) {
    resArray.push(item);
  }
  console.log(resArray);
}

workbooksList().catch(console.error);
