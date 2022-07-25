const { TimeSeriesInsightsClient } = require("@azure/arm-timeseriesinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates a Private Endpoint connection of the environment in the given resource group.
 *
 * @summary Updates a Private Endpoint connection of the environment in the given resource group.
 * x-ms-original-file: specification/timeseriesinsights/resource-manager/Microsoft.TimeSeriesInsights/preview/2021-03-31-preview/examples/PrivateEndpointConnectionUpdate.json
 */
async function privateEndpointConnectionUpdate() {
  const subscriptionId = "mySubscriptionId";
  const resourceGroupName = "myResourceGroup";
  const environmentName = "myEnvironment";
  const privateEndpointConnectionName = "myPrivateEndpointConnectionName";
  const privateEndpointConnection = {
    privateLinkServiceConnectionState: {
      description: "Rejected for some reason",
      status: "Rejected",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new TimeSeriesInsightsClient(credential, subscriptionId);
  const result = await client.privateEndpointConnections.createOrUpdate(
    resourceGroupName,
    environmentName,
    privateEndpointConnectionName,
    privateEndpointConnection
  );
  console.log(result);
}

privateEndpointConnectionUpdate().catch(console.error);
