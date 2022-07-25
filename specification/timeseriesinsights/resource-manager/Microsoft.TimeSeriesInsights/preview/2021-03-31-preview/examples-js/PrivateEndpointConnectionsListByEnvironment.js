const { TimeSeriesInsightsClient } = require("@azure/arm-timeseriesinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a list of all private endpoint connections in the given environment.
 *
 * @summary Gets a list of all private endpoint connections in the given environment.
 * x-ms-original-file: specification/timeseriesinsights/resource-manager/Microsoft.TimeSeriesInsights/preview/2021-03-31-preview/examples/PrivateEndpointConnectionsListByEnvironment.json
 */
async function listPrivateEndpointConnectionsByService() {
  const subscriptionId = "mySubscriptionId";
  const resourceGroupName = "myResourceGroup";
  const environmentName = "myEnvironment";
  const credential = new DefaultAzureCredential();
  const client = new TimeSeriesInsightsClient(credential, subscriptionId);
  const result = await client.privateEndpointConnections.listByEnvironment(
    resourceGroupName,
    environmentName
  );
  console.log(result);
}

listPrivateEndpointConnectionsByService().catch(console.error);
