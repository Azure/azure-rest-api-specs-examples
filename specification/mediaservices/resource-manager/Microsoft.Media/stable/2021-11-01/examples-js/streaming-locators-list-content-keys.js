const { AzureMediaServices } = require("@azure/arm-mediaservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List Content Keys used by this Streaming Locator
 *
 * @summary List Content Keys used by this Streaming Locator
 * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/streaming-locators-list-content-keys.json
 */
async function listContentKeys() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "contoso";
  const accountName = "contosomedia";
  const streamingLocatorName = "secureStreamingLocator";
  const credential = new DefaultAzureCredential();
  const client = new AzureMediaServices(credential, subscriptionId);
  const result = await client.streamingLocators.listContentKeys(
    resourceGroupName,
    accountName,
    streamingLocatorName
  );
  console.log(result);
}

listContentKeys().catch(console.error);
