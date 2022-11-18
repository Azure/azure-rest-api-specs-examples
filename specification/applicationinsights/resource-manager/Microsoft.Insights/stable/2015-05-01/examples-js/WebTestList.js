const { ApplicationInsightsManagementClient } = require("@azure/arm-appinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get all Application Insights web test alerts definitions within a subscription.
 *
 * @summary Get all Application Insights web test alerts definitions within a subscription.
 * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/WebTestList.json
 */
async function webTestList() {
  const subscriptionId = "subid";
  const credential = new DefaultAzureCredential();
  const client = new ApplicationInsightsManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.webTests.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

webTestList().catch(console.error);
