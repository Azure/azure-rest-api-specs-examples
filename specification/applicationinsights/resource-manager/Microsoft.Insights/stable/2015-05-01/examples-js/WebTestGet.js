const { ApplicationInsightsManagementClient } = require("@azure/arm-appinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a specific Application Insights web test definition.
 *
 * @summary Get a specific Application Insights web test definition.
 * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/WebTestGet.json
 */
async function webTestGet() {
  const subscriptionId = "subid";
  const resourceGroupName = "my-resource-group";
  const webTestName = "my-webtest-01-mywebservice";
  const credential = new DefaultAzureCredential();
  const client = new ApplicationInsightsManagementClient(credential, subscriptionId);
  const result = await client.webTests.get(resourceGroupName, webTestName);
  console.log(result);
}

webTestGet().catch(console.error);
