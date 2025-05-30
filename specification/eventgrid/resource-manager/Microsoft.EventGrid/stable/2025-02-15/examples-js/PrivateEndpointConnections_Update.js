const { EventGridManagementClient } = require("@azure/arm-eventgrid");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Update a specific private endpoint connection under a topic, domain or partner namespace.
 *
 * @summary Update a specific private endpoint connection under a topic, domain or partner namespace.
 * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2025-02-15/examples/PrivateEndpointConnections_Update.json
 */
async function privateEndpointConnectionsUpdate() {
  const subscriptionId =
    process.env["EVENTGRID_SUBSCRIPTION_ID"] || "5b4b650e-28b9-4790-b3ab-ddbd88d727c4";
  const resourceGroupName = process.env["EVENTGRID_RESOURCE_GROUP"] || "examplerg";
  const parentType = "topics";
  const parentName = "exampletopic1";
  const privateEndpointConnectionName = "BMTPE5.8A30D251-4C61-489D-A1AA-B37C4A329B8B";
  const privateEndpointConnection = {
    privateLinkServiceConnectionState: {
      description: "approving connection",
      actionsRequired: "None",
      status: "Approved",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new EventGridManagementClient(credential, subscriptionId);
  const result = await client.privateEndpointConnections.beginUpdateAndWait(
    resourceGroupName,
    parentType,
    parentName,
    privateEndpointConnectionName,
    privateEndpointConnection,
  );
  console.log(result);
}
