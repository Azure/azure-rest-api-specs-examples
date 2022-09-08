const { AzureMediaServices } = require("@azure/arm-mediaservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create a Streaming Policy in the Media Services account
 *
 * @summary Create a Streaming Policy in the Media Services account
 * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2022-08-01/examples/streaming-policies-create-envelopeEncryption-only.json
 */
async function createsAStreamingPolicyWithEnvelopeEncryptionOnly() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "contoso";
  const accountName = "contosomedia";
  const streamingPolicyName = "UserCreatedSecureStreamingPolicyWithEnvelopeEncryptionOnly";
  const parameters = {
    defaultContentKeyPolicyName: "PolicyWithClearKeyOptionAndTokenRestriction",
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

createsAStreamingPolicyWithEnvelopeEncryptionOnly().catch(console.error);
