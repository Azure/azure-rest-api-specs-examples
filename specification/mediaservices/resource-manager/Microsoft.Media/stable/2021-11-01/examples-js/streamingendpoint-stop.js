const { AzureMediaServices } = require("@azure/arm-mediaservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Stops an existing streaming endpoint.
 *
 * @summary Stops an existing streaming endpoint.
 * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/streamingendpoint-stop.json
 */
async function stopAStreamingEndpoint() {
  const subscriptionId = "0a6ec948-5a62-437d-b9df-934dc7c1b722";
  const resourceGroupName = "mediaresources";
  const accountName = "slitestmedia10";
  const streamingEndpointName = "myStreamingEndpoint1";
  const credential = new DefaultAzureCredential();
  const client = new AzureMediaServices(credential, subscriptionId);
  const result = await client.streamingEndpoints.beginStopAndWait(
    resourceGroupName,
    accountName,
    streamingEndpointName
  );
  console.log(result);
}

stopAStreamingEndpoint().catch(console.error);
