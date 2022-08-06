const { AzureMediaServices } = require("@azure/arm-mediaservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all the live events in the account.
 *
 * @summary Lists all the live events in the account.
 * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/liveevent-list-all.json
 */
async function listAllLiveEvents() {
  const subscriptionId = "0a6ec948-5a62-437d-b9df-934dc7c1b722";
  const resourceGroupName = "mediaresources";
  const accountName = "slitestmedia10";
  const credential = new DefaultAzureCredential();
  const client = new AzureMediaServices(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.liveEvents.list(resourceGroupName, accountName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listAllLiveEvents().catch(console.error);
