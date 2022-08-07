const { AzureMediaServices } = require("@azure/arm-mediaservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a Streaming Policy in the Media Services account
 *
 * @summary Deletes a Streaming Policy in the Media Services account
 * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/streaming-policies-delete.json
 */
async function deleteAStreamingPolicy() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "contoso";
  const accountName = "contosomedia";
  const streamingPolicyName = "secureStreamingPolicyWithCommonEncryptionCbcsOnly";
  const credential = new DefaultAzureCredential();
  const client = new AzureMediaServices(credential, subscriptionId);
  const result = await client.streamingPolicies.delete(
    resourceGroupName,
    accountName,
    streamingPolicyName
  );
  console.log(result);
}

deleteAStreamingPolicy().catch(console.error);
