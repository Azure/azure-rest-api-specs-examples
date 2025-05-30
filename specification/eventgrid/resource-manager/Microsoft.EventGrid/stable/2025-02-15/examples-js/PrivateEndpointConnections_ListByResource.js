const { EventGridManagementClient } = require("@azure/arm-eventgrid");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Get all private endpoint connections under a topic, domain, or partner namespace or namespace.
 *
 * @summary Get all private endpoint connections under a topic, domain, or partner namespace or namespace.
 * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2025-02-15/examples/PrivateEndpointConnections_ListByResource.json
 */
async function privateEndpointConnectionsListByResource() {
  const subscriptionId =
    process.env["EVENTGRID_SUBSCRIPTION_ID"] || "5b4b650e-28b9-4790-b3ab-ddbd88d727c4";
  const resourceGroupName = process.env["EVENTGRID_RESOURCE_GROUP"] || "examplerg";
  const parentType = "topics";
  const parentName = "exampletopic1";
  const credential = new DefaultAzureCredential();
  const client = new EventGridManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.privateEndpointConnections.listByResource(
    resourceGroupName,
    parentType,
    parentName,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
