const { AzureMediaServices } = require("@azure/arm-mediaservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates a existing streaming endpoint.
 *
 * @summary Updates a existing streaming endpoint.
 * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2022-08-01/examples/streamingendpoint-update.json
 */
async function updateAStreamingEndpoint() {
  const subscriptionId = "0a6ec948-5a62-437d-b9df-934dc7c1b722";
  const resourceGroupName = "mediaresources";
  const accountName = "slitestmedia10";
  const streamingEndpointName = "myStreamingEndpoint1";
  const parameters = {
    description: "test event 2",
    availabilitySetName: "availableset",
    location: "West US",
    scaleUnits: 5,
    tags: { tag3: "value3", tag5: "value5" },
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureMediaServices(credential, subscriptionId);
  const result = await client.streamingEndpoints.beginUpdateAndWait(
    resourceGroupName,
    accountName,
    streamingEndpointName,
    parameters
  );
  console.log(result);
}

updateAStreamingEndpoint().catch(console.error);
