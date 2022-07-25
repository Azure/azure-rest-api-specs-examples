const { TimeSeriesInsightsClient } = require("@azure/arm-timeseriesinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the details of the private endpoint connection of the environment in the given resource group.
 *
 * @summary Gets the details of the private endpoint connection of the environment in the given resource group.
 * x-ms-original-file: specification/timeseriesinsights/resource-manager/Microsoft.TimeSeriesInsights/preview/2021-03-31-preview/examples/PrivateEndpointConnectionGet.json
 */
async function privateEndpointConnectionGet() {
  const subscriptionId = "mySubscriptionId";
  const resourceGroupName = "myResourceGroup";
  const environmentName = "myEnvironment";
  const privateEndpointConnectionName = "myPrivateEndpointConnectionName";
  const credential = new DefaultAzureCredential();
  const client = new TimeSeriesInsightsClient(credential, subscriptionId);
  const result = await client.privateEndpointConnections.get(
    resourceGroupName,
    environmentName,
    privateEndpointConnectionName
  );
  console.log(result);
}

privateEndpointConnectionGet().catch(console.error);
