const { AzureMediaServices } = require("@azure/arm-mediaservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create a Streaming Policy in the Media Services account
 *
 * @summary Create a Streaming Policy in the Media Services account
 * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/streaming-policies-create-commonEncryptionCenc-only.json
 */
async function createsAStreamingPolicyWithCommonEncryptionCencOnly() {
  const subscriptionId =
    process.env["MEDIASERVICES_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["MEDIASERVICES_RESOURCE_GROUP"] || "contoso";
  const accountName = "contosomedia";
  const streamingPolicyName = "UserCreatedSecureStreamingPolicyWithCommonEncryptionCencOnly";
  const parameters = {
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
    defaultContentKeyPolicyName: "PolicyWithPlayReadyOptionAndOpenRestriction",
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
