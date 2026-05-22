const { ApplicationInsightsManagementClient } = require("@azure/arm-appinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to get all Workbooks defined within a specified subscription and category.
 *
 * @summary get all Workbooks defined within a specified subscription and category.
 * x-ms-original-file: 2023-06-01/WorkbooksList2.json
 */
async function workbooksList2() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "6b643656-33eb-422f-aee8-3ac145d124af";
  const client = new ApplicationInsightsManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.workbooks.listBySubscription("workbook")) {
    resArray.push(item);
  }

  console.log(resArray);
}
