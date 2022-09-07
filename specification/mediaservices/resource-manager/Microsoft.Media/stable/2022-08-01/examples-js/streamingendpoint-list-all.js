const { AzureMediaServices } = require("@azure/arm-mediaservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the streaming endpoints in the account.
 *
 * @summary Lists the streaming endpoints in the account.
 * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2022-08-01/examples/streamingendpoint-list-all.json
 */
async function listAllStreamingEndpoints() {
  const subscriptionId = "0a6ec948-5a62-437d-b9df-934dc7c1b722";
  const resourceGroupName = "mediaresources";
  const accountName = "slitestmedia10";
  const credential = new DefaultAzureCredential();
  const client = new AzureMediaServices(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.streamingEndpoints.list(resourceGroupName, accountName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listAllStreamingEndpoints().catch(console.error);
