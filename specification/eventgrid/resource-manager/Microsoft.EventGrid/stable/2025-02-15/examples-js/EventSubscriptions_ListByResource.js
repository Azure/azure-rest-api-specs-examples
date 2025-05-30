const { EventGridManagementClient } = require("@azure/arm-eventgrid");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to List all event subscriptions that have been created for a specific resource.
 *
 * @summary List all event subscriptions that have been created for a specific resource.
 * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2025-02-15/examples/EventSubscriptions_ListByResource.json
 */
async function eventSubscriptionsListByResource() {
  const subscriptionId =
    process.env["EVENTGRID_SUBSCRIPTION_ID"] || "5b4b650e-28b9-4790-b3ab-ddbd88d727c4";
  const resourceGroupName = process.env["EVENTGRID_RESOURCE_GROUP"] || "examplerg";
  const providerNamespace = "Microsoft.EventGrid";
  const resourceTypeName = "topics";
  const resourceName = "exampletopic2";
  const credential = new DefaultAzureCredential();
  const client = new EventGridManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.eventSubscriptions.listByResource(
    resourceGroupName,
    providerNamespace,
    resourceTypeName,
    resourceName,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
