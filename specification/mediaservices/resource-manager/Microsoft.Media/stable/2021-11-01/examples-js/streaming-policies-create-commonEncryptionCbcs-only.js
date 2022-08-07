const { AzureMediaServices } = require("@azure/arm-mediaservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create a Streaming Policy in the Media Services account
 *
 * @summary Create a Streaming Policy in the Media Services account
 * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/streaming-policies-create-commonEncryptionCbcs-only.json
 */
async function createsAStreamingPolicyWithCommonEncryptionCbcsOnly() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "contoso";
  const accountName = "contosomedia";
  const streamingPolicyName = "UserCreatedSecureStreamingPolicyWithCommonEncryptionCbcsOnly";
  const parameters = {
    commonEncryptionCbcs: {
      contentKeys: { defaultKey: { label: "cbcsDefaultKey" } },
      drm: {
        fairPlay: {
          allowPersistentLicense: true,
          customLicenseAcquisitionUrlTemplate:
            "https://contoso.com/{AssetAlternativeId}/fairplay/{ContentKeyId}",
        },
      },
      enabledProtocols: {
        dash: false,
        download: false,
        hls: true,
        smoothStreaming: false,
      },
    },
    defaultContentKeyPolicyName: "PolicyWithMultipleOptions",
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureMediaServices(credential, subscriptionId);
  const result = await client.streamingPolicies.create(
    resourceGroupName,
    accountName,
    streamingPolicyName,
    parameters
  );
  console.log(result);
}

createsAStreamingPolicyWithCommonEncryptionCbcsOnly().catch(console.error);
