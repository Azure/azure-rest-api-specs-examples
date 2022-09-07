const { AzureMediaServices } = require("@azure/arm-mediaservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a live output. Deleting a live output does not delete the asset the live output is writing to.
 *
 * @summary Deletes a live output. Deleting a live output does not delete the asset the live output is writing to.
 * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2022-08-01/examples/liveoutput-delete.json
 */
async function deleteALiveOutput() {
  const subscriptionId = "0a6ec948-5a62-437d-b9df-934dc7c1b722";
  const resourceGroupName = "mediaresources";
  const accountName = "slitestmedia10";
  const liveEventName = "myLiveEvent1";
  const liveOutputName = "myLiveOutput1";
  const credential = new DefaultAzureCredential();
  const client = new AzureMediaServices(credential, subscriptionId);
  const result = await client.liveOutputs.beginDeleteAndWait(
    resourceGroupName,
    accountName,
    liveEventName,
    liveOutputName
  );
  console.log(result);
}

deleteALiveOutput().catch(console.error);
