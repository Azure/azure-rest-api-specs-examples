const { AzureMediaServices } = require("@azure/arm-mediaservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create a Streaming Policy in the Media Services account
 *
 * @summary Create a Streaming Policy in the Media Services account
 * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2022-08-01/examples/streaming-policies-create-secure-streaming.json
 */
async function createsAStreamingPolicyWithSecureStreaming() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "contoso";
  const accountName = "contosomedia";
  const streamingPolicyName = "UserCreatedSecureStreamingPolicy";
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
    commonEncryptionCenc: {
      clearTracks: [
        {
          trackSelections: [{ operation: "Equal", property: "FourCC", value: "hev1" }],
        },
      ],
      contentKeys: { defaultKey: { label: "cencDefaultKey" } },
      drm: {
        playReady: {
          customLicenseAcquisitionUrlTemplate:
            "https://contoso.com/{AssetAlternativeId}/playready/{ContentKeyId}",
          playReadyCustomAttributes: "PlayReady CustomAttributes",
        },
        widevine: {
          customLicenseAcquisitionUrlTemplate:
            "https://contoso.com/{AssetAlternativeId}/widevine/{ContentKeyId",
        },
      },
      enabledProtocols: {
        dash: true,
        download: false,
        hls: false,
        smoothStreaming: true,
      },
    },
    defaultContentKeyPolicyName: "PolicyWithMultipleOptions",
    envelopeEncryption: {
      contentKeys: { defaultKey: { label: "aesDefaultKey" } },
      customKeyAcquisitionUrlTemplate:
        "https://contoso.com/{AssetAlternativeId}/envelope/{ContentKeyId}",
      enabledProtocols: {
        dash: true,
        download: false,
        hls: true,
        smoothStreaming: true,
      },
    },
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

createsAStreamingPolicyWithSecureStreaming().catch(console.error);
