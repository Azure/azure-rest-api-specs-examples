const { EventGridManagementClient } = require("@azure/arm-eventgrid");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List all event subscriptions from the given location under a specific Azure subscription and resource group.
 *
 * @summary List all event subscriptions from the given location under a specific Azure subscription and resource group.
 * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2022-06-15/examples/EventSubscriptions_ListRegionalByResourceGroup.json
 */
async function eventSubscriptionsListRegionalByResourceGroup() {
  const subscriptionId = "5b4b650e-28b9-4790-b3ab-ddbd88d727c4";
  const resourceGroupName = "examplerg";
  const location = "westus2";
  const credential = new DefaultAzureCredential();
  const client = new EventGridManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.eventSubscriptions.listRegionalByResourceGroup(
    resourceGroupName,
    location
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

eventSubscriptionsListRegionalByResourceGroup().catch(console.error);
