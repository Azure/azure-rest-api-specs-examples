const { AzureMediaServices } = require("@azure/arm-mediaservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List Paths supported by this Streaming Locator
 *
 * @summary List Paths supported by this Streaming Locator
 * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/streaming-locators-list-paths-streaming-only.json
 */
async function listPathsWhichHasStreamingPathsOnly() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "contoso";
  const accountName = "contosomedia";
  const streamingLocatorName = "secureStreamingLocator";
  const credential = new DefaultAzureCredential();
  const client = new AzureMediaServices(credential, subscriptionId);
  const result = await client.streamingLocators.listPaths(
    resourceGroupName,
    accountName,
    streamingLocatorName
  );
  console.log(result);
}

listPathsWhichHasStreamingPathsOnly().catch(console.error);
