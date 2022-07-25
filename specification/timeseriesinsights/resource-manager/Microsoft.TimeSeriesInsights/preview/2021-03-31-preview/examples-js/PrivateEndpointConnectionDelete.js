const { TimeSeriesInsightsClient } = require("@azure/arm-timeseriesinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Disconnects the private endpoint connection and deletes it from the environment.
 *
 * @summary Disconnects the private endpoint connection and deletes it from the environment.
 * x-ms-original-file: specification/timeseriesinsights/resource-manager/Microsoft.TimeSeriesInsights/preview/2021-03-31-preview/examples/PrivateEndpointConnectionDelete.json
 */
async function privateEndpointConnectionDelete() {
  const subscriptionId = "mySubscriptionId";
  const resourceGroupName = "myResourceGroup";
  const environmentName = "myEnvironment";
  const privateEndpointConnectionName = "myPrivateEndpointConnectionName";
  const credential = new DefaultAzureCredential();
  const client = new TimeSeriesInsightsClient(credential, subscriptionId);
  const result = await client.privateEndpointConnections.delete(
    resourceGroupName,
    environmentName,
    privateEndpointConnectionName
  );
  console.log(result);
}

privateEndpointConnectionDelete().catch(console.error);
