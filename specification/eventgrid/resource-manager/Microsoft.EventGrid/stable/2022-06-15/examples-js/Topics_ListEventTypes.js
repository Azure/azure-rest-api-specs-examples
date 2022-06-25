const { EventGridManagementClient } = require("@azure/arm-eventgrid");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List event types for a topic.
 *
 * @summary List event types for a topic.
 * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2022-06-15/examples/Topics_ListEventTypes.json
 */
async function topicsListEventTypes() {
  const subscriptionId = "5b4b650e-28b9-4790-b3ab-ddbd88d727c4";
  const resourceGroupName = "examplerg";
  const providerNamespace = "Microsoft.Storage";
  const resourceTypeName = "storageAccounts";
  const resourceName = "ExampleStorageAccount";
  const credential = new DefaultAzureCredential();
  const client = new EventGridManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.topics.listEventTypes(
    resourceGroupName,
    providerNamespace,
    resourceTypeName,
    resourceName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

topicsListEventTypes().catch(console.error);
