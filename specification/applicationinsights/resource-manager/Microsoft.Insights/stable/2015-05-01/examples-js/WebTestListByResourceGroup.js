const { ApplicationInsightsManagementClient } = require("@azure/arm-appinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get all Application Insights web tests defined within a specified resource group.
 *
 * @summary Get all Application Insights web tests defined within a specified resource group.
 * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/WebTestListByResourceGroup.json
 */
async function webTestListByResourceGroup() {
  const subscriptionId = "subid";
  const resourceGroupName = "my-resource-group";
  const credential = new DefaultAzureCredential();
  const client = new ApplicationInsightsManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.webTests.listByResourceGroup(resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

webTestListByResourceGroup().catch(console.error);
