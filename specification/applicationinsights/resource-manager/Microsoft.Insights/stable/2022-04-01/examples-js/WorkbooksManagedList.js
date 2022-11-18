const { ApplicationInsightsManagementClient } = require("@azure/arm-appinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get all Workbooks defined within a specified resource group and category.
 *
 * @summary Get all Workbooks defined within a specified resource group and category.
 * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2022-04-01/examples/WorkbooksManagedList.json
 */
async function workbooksManagedList() {
  const subscriptionId = "6b643656-33eb-422f-aee8-3ac119r124af";
  const resourceGroupName = "my-resource-group";
  const category = "workbook";
  const sourceId =
    "/subscriptions/6b643656-33eb-422f-aee8-3ac119r124af/resourceGroups/my-resource-group/providers/Microsoft.Web/sites/MyApp";
  const options = { sourceId };
  const credential = new DefaultAzureCredential();
  const client = new ApplicationInsightsManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.workbooks.listByResourceGroup(
    resourceGroupName,
    category,
    options
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

workbooksManagedList().catch(console.error);
