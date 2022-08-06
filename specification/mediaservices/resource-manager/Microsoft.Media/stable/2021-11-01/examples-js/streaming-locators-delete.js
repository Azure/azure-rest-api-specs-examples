const { AzureMediaServices } = require("@azure/arm-mediaservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a Streaming Locator in the Media Services account
 *
 * @summary Deletes a Streaming Locator in the Media Services account
 * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/streaming-locators-delete.json
 */
async function deleteAStreamingLocator() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "contoso";
  const accountName = "contosomedia";
  const streamingLocatorName = "clearStreamingLocator";
  const credential = new DefaultAzureCredential();
  const client = new AzureMediaServices(credential, subscriptionId);
  const result = await client.streamingLocators.delete(
    resourceGroupName,
    accountName,
    streamingLocatorName
  );
  console.log(result);
}

deleteAStreamingLocator().catch(console.error);
