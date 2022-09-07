const { AzureMediaServices } = require("@azure/arm-mediaservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates settings on an existing live event.
 *
 * @summary Updates settings on an existing live event.
 * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2022-08-01/examples/liveevent-update.json
 */
async function updateALiveEvent() {
  const subscriptionId = "0a6ec948-5a62-437d-b9df-934dc7c1b722";
  const resourceGroupName = "mediaresources";
  const accountName = "slitestmedia10";
  const liveEventName = "myLiveEvent1";
  const parameters = {
    description: "test event updated",
    input: {
      accessControl: {
        ip: { allow: [{ name: "AllowOne", address: "192.1.1.0" }] },
      },
      keyFrameIntervalDuration: "PT6S",
      streamingProtocol: "FragmentedMP4",
    },
    location: "West US",
    preview: {
      accessControl: {
        ip: { allow: [{ name: "AllowOne", address: "192.1.1.0" }] },
      },
    },
    tags: { tag1: "value1", tag2: "value2", tag3: "value3" },
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureMediaServices(credential, subscriptionId);
  const result = await client.liveEvents.beginUpdateAndWait(
    resourceGroupName,
    accountName,
    liveEventName,
    parameters
  );
  console.log(result);
}

updateALiveEvent().catch(console.error);
