const { AzureMediaServices } = require("@azure/arm-mediaservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the live outputs of a live event.
 *
 * @summary Lists the live outputs of a live event.
 * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2022-08-01/examples/liveoutput-list-all.json
 */
async function listAllLiveOutputs() {
  const subscriptionId = "0a6ec948-5a62-437d-b9df-934dc7c1b722";
  const resourceGroupName = "mediaresources";
  const accountName = "slitestmedia10";
  const liveEventName = "myLiveEvent1";
  const credential = new DefaultAzureCredential();
  const client = new AzureMediaServices(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.liveOutputs.list(resourceGroupName, accountName, liveEventName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listAllLiveOutputs().catch(console.error);
