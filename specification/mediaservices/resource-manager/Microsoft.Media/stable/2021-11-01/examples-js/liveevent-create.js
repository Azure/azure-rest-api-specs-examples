const { AzureMediaServices } = require("@azure/arm-mediaservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a new live event.
 *
 * @summary Creates a new live event.
 * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/liveevent-create.json
 */
async function createALiveEvent() {
  const subscriptionId = "0a6ec948-5a62-437d-b9df-934dc7c1b722";
  const resourceGroupName = "mediaresources";
  const accountName = "slitestmedia10";
  const liveEventName = "myLiveEvent1";
  const parameters = {
    description: "test event 1",
    input: {
      accessControl: {
        ip: {
          allow: [{ name: "AllowAll", address: "0.0.0.0", subnetPrefixLength: 0 }],
        },
      },
      keyFrameIntervalDuration: "PT6S",
      streamingProtocol: "RTMP",
    },
    location: "West US",
    preview: {
      accessControl: {
        ip: {
          allow: [{ name: "AllowAll", address: "0.0.0.0", subnetPrefixLength: 0 }],
        },
      },
    },
    tags: { tag1: "value1", tag2: "value2" },
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureMediaServices(credential, subscriptionId);
  const result = await client.liveEvents.beginCreateAndWait(
    resourceGroupName,
    accountName,
    liveEventName,
    parameters
  );
  console.log(result);
}

createALiveEvent().catch(console.error);
