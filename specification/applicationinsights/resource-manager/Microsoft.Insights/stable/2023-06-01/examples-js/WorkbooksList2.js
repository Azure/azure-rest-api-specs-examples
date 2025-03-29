const { ApplicationInsightsManagementClient } = require("@azure/arm-appinsights");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Get all Workbooks defined within a specified subscription and category.
 *
 * @summary Get all Workbooks defined within a specified subscription and category.
 * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2023-06-01/examples/WorkbooksList2.json
 */
async function workbooksList2() {
  const subscriptionId =
    process.env["APPLICATIONINSIGHTS_SUBSCRIPTION_ID"] || "6b643656-33eb-422f-aee8-3ac145d124af";
  const category = "workbook";
  const credential = new DefaultAzureCredential();
  const client = new ApplicationInsightsManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.workbooks.listBySubscription(category)) {
    resArray.push(item);
  }
  console.log(resArray);
}
