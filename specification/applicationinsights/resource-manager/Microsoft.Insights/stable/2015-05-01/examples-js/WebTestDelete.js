const { ApplicationInsightsManagementClient } = require("@azure/arm-appinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes an Application Insights web test.
 *
 * @summary Deletes an Application Insights web test.
 * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/WebTestDelete.json
 */
async function webTestDelete() {
  const subscriptionId = "subid";
  const resourceGroupName = "my-resource-group";
  const webTestName = "my-webtest-01-mywebservice";
  const credential = new DefaultAzureCredential();
  const client = new ApplicationInsightsManagementClient(credential, subscriptionId);
  const result = await client.webTests.delete(resourceGroupName, webTestName);
  console.log(result);
}

webTestDelete().catch(console.error);
