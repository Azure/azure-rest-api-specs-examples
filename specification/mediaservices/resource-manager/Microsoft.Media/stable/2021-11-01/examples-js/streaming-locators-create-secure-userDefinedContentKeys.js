const { AzureMediaServices } = require("@azure/arm-mediaservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create a Streaming Locator in the Media Services account
 *
 * @summary Create a Streaming Locator in the Media Services account
 * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/streaming-locators-create-secure-userDefinedContentKeys.json
 */
async function createsAStreamingLocatorWithUserDefinedContentKeys() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "contoso";
  const accountName = "contosomedia";
  const streamingLocatorName = "UserCreatedSecureStreamingLocatorWithUserDefinedContentKeys";
  const parameters = {
    assetName: "ClimbingMountRainier",
    contentKeys: [
      {
        id: "60000000-0000-0000-0000-000000000001",
        labelReferenceInStreamingPolicy: "aesDefaultKey",
        value: "1UqLohAfWsEGkULYxHjYZg==",
      },
      {
        id: "60000000-0000-0000-0000-000000000004",
        labelReferenceInStreamingPolicy: "cencDefaultKey",
        value: "4UqLohAfWsEGkULYxHjYZg==",
      },
      {
        id: "60000000-0000-0000-0000-000000000007",
        labelReferenceInStreamingPolicy: "cbcsDefaultKey",
        value: "7UqLohAfWsEGkULYxHjYZg==",
      },
    ],
    streamingLocatorId: "90000000-0000-0000-0000-00000000000A",
    streamingPolicyName: "secureStreamingPolicy",
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

createsAStreamingLocatorWithUserDefinedContentKeys().catch(console.error);
