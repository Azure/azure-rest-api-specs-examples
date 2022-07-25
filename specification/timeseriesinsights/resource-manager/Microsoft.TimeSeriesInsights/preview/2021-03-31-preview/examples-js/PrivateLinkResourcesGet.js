const { TimeSeriesInsightsClient } = require("@azure/arm-timeseriesinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a list of all supported private link resource types for the given environment.
 *
 * @summary Gets a list of all supported private link resource types for the given environment.
 * x-ms-original-file: specification/timeseriesinsights/resource-manager/Microsoft.TimeSeriesInsights/preview/2021-03-31-preview/examples/PrivateLinkResourcesGet.json
 */
async function listSupportedPrivateLinkResources() {
  const subscriptionId = "mySubscriptionId";
  const resourceGroupName = "myResourceGroup";
  const environmentName = "myEnvironment";
  const credential = new DefaultAzureCredential();
  const client = new TimeSeriesInsightsClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.privateLinkResources.listSupported(
    resourceGroupName,
    environmentName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

listSupportedPrivateLinkResources().catch(console.error);
