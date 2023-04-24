const { EventHubManagementClient } = require("@azure/arm-eventhub");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates PrivateEndpointConnections of service namespace.
 *
 * @summary Creates or updates PrivateEndpointConnections of service namespace.
 * x-ms-original-file: specification/eventhub/resource-manager/Microsoft.EventHub/preview/2022-10-01-preview/examples/NameSpaces/PrivateEndPointConnectionCreate.json
 */
async function nameSpacePrivateEndPointConnectionCreate() {
  const subscriptionId = process.env["EVENTHUB_SUBSCRIPTION_ID"] || "subID";
  const resourceGroupName = process.env["EVENTHUB_RESOURCE_GROUP"] || "ArunMonocle";
  const namespaceName = "sdk-Namespace-2924";
  const privateEndpointConnectionName = "privateEndpointConnectionName";
  const parameters = {
    privateEndpoint: {
      id: "/subscriptions/dbedb4e0-40e6-4145-81f3-f1314c150774/resourceGroups/SDK-EventHub-8396/providers/Microsoft.Network/privateEndpoints/sdk-Namespace-2847",
    },
    privateLinkServiceConnectionState: {
      description: "testing",
      status: "Rejected",
    },
    provisioningState: "Succeeded",
  };
  const credential = new DefaultAzureCredential();
  const client = new EventHubManagementClient(credential, subscriptionId);
  const result = await client.privateEndpointConnections.createOrUpdate(
    resourceGroupName,
    namespaceName,
    privateEndpointConnectionName,
    parameters
  );
  console.log(result);
}
