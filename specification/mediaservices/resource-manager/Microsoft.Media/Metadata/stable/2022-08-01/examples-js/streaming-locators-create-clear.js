const { AzureMediaServices } = require("@azure/arm-mediaservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create a Streaming Locator in the Media Services account
 *
 * @summary Create a Streaming Locator in the Media Services account
 * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/streaming-locators-create-clear.json
 */
async function createsAStreamingLocatorWithClearStreaming() {
  const subscriptionId =
    process.env["MEDIASERVICES_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["MEDIASERVICES_RESOURCE_GROUP"] || "contoso";
  const accountName = "contosomedia";
  const streamingLocatorName = "UserCreatedClearStreamingLocator";
  const parameters = {
    assetName: "ClimbingMountRainier",
    streamingPolicyName: "clearStreamingPolicy",
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureMediaServices(credential, subscriptionId);
  const result = await client.streamingLocators.create(
    resourceGroupName,
    accountName,
    streamingLocatorName,
    parameters
  );
  console.log(result);
}
