const { CustomerInsightsManagementClient } = require("@azure/arm-customerinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all the hubs in a resource group.
 *
 * @summary Gets all the hubs in a resource group.
 * x-ms-original-file: specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/HubsListByResourceGroup.json
 */
async function hubsListByResourceGroup() {
  const subscriptionId = "subid";
  const resourceGroupName = "TestHubRG";
  const credential = new DefaultAzureCredential();
  const client = new CustomerInsightsManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.hubs.listByResourceGroup(resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

hubsListByResourceGroup().catch(console.error);
