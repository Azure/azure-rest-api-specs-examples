const { AzureMediaServices } = require("@azure/arm-mediaservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get asset track operation status.
 *
 * @summary Get asset track operation status.
 * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/asset-tracks-operation-status-by-id-terminal-state-failed.json
 */
async function getStatusOfAsynchronousOperationWhenItIsCompletedWithError() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "contoso";
  const accountName = "contosomedia";
  const assetName = "ClimbingMountRainer";
  const trackName = "text1";
  const operationId = "86835197-3b47-402e-b313-70b82eaba296";
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

getStatusOfAsynchronousOperationWhenItIsCompletedWithError().catch(console.error);
