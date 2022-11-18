const { ApplicationInsightsManagementClient } = require("@azure/arm-appinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get all Application Insights web tests defined for the specified component.
 *
 * @summary Get all Application Insights web tests defined for the specified component.
 * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/WebTestListByComponent.json
 */
async function webTestListByComponent() {
  const subscriptionId = "subid";
  const componentName = "my-component";
  const resourceGroupName = "my-resource-group";
  const credential = new DefaultAzureCredential();
  const client = new ApplicationInsightsManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.webTests.listByComponent(componentName, resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

webTestListByComponent().catch(console.error);
