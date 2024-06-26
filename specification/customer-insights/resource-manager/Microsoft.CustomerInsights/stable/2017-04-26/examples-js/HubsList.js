const { CustomerInsightsManagementClient } = require("@azure/arm-customerinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all hubs in the specified subscription.
 *
 * @summary Gets all hubs in the specified subscription.
 * x-ms-original-file: specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/HubsList.json
 */
async function hubsList() {
  const subscriptionId = "subid";
  const credential = new DefaultAzureCredential();
  const client = new CustomerInsightsManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.hubs.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

hubsList().catch(console.error);
