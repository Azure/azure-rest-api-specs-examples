const { AzureMediaServices } = require("@azure/arm-mediaservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Resets an existing live event. All live outputs for the live event are deleted and the live event is stopped and will be started again. All assets used by the live outputs and streaming locators created on these assets are unaffected.
 *
 * @summary Resets an existing live event. All live outputs for the live event are deleted and the live event is stopped and will be started again. All assets used by the live outputs and streaming locators created on these assets are unaffected.
 * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/liveevent-reset.json
 */
async function resetALiveEvent() {
  const subscriptionId = "0a6ec948-5a62-437d-b9df-934dc7c1b722";
  const resourceGroupName = "mediaresources";
  const accountName = "slitestmedia10";
  const liveEventName = "myLiveEvent1";
  const credential = new DefaultAzureCredential();
  const client = new AzureMediaServices(credential, subscriptionId);
  const result = await client.liveEvents.beginResetAndWait(
    resourceGroupName,
    accountName,
    liveEventName
  );
  console.log(result);
}

resetALiveEvent().catch(console.error);
