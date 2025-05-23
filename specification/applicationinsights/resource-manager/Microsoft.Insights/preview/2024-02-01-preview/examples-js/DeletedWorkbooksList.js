const { ApplicationInsightsManagementClient } = require("@azure/arm-appinsights");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Get all recently deleted Workbooks in a specified subscription.
 *
 * @summary Get all recently deleted Workbooks in a specified subscription.
 * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/preview/2024-02-01-preview/examples/DeletedWorkbooksList.json
 */
async function workbooksListSub() {
  const subscriptionId =
    process.env["APPLICATIONINSIGHTS_SUBSCRIPTION_ID"] || "00000000-1111-2222-3333-444444444444";
  const credential = new DefaultAzureCredential();
  const client = new ApplicationInsightsManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.deletedWorkbooks.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}
