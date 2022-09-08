const { AzureMediaServices } = require("@azure/arm-mediaservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get asset track operation status.
 *
 * @summary Get asset track operation status.
 * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2022-08-01/examples/asset-tracks-operation-status-by-id-terminal-state.json
 */
async function getStatusOfAsynchronousOperationWhenItIsCompleted() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "contoso";
  const accountName = "contosomedia";
  const assetName = "ClimbingMountRainer";
  const trackName = "text1";
  const operationId = "e78f8d40-7aaa-4f2f-8ae6-73987e7c5a08";
  const credential = new DefaultAzureCredential();
  const client = new AzureMediaServices(credential, subscriptionId);
  const result = await client.operationStatuses.get(
    resourceGroupName,
    accountName,
    assetName,
    trackName,
    operationId
  );
  console.log(result);
}

getStatusOfAsynchronousOperationWhenItIsCompleted().catch(console.error);
