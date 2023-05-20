const { EventGridManagementClient } = require("@azure/arm-eventgrid");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get all private endpoint connections under a topic, domain, or partner namespace or namespace.
 *
 * @summary Get all private endpoint connections under a topic, domain, or partner namespace or namespace.
 * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-06-01-preview/examples/PrivateEndpointConnections_ListByResource.json
 */
async function privateEndpointConnectionsListByResource() {
  const subscriptionId =
    process.env["EVENTGRID_SUBSCRIPTION_ID"] || "8f6b6269-84f2-4d09-9e31-1127efcd1e40";
  const resourceGroupName = process.env["EVENTGRID_RESOURCE_GROUP"] || "examplerg";
  const parentType = "topics";
  const parentName = "exampletopic1";
  const credential = new DefaultAzureCredential();
  const client = new EventGridManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.privateEndpointConnections.listByResource(
    resourceGroupName,
    parentType,
    parentName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
