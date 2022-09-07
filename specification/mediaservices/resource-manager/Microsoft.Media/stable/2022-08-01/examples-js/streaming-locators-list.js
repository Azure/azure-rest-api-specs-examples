const { AzureMediaServices } = require("@azure/arm-mediaservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the Streaming Locators in the account
 *
 * @summary Lists the Streaming Locators in the account
 * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2022-08-01/examples/streaming-locators-list.json
 */
async function listsStreamingLocators() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "contoso";
  const accountName = "contosomedia";
  const credential = new DefaultAzureCredential();
  const client = new AzureMediaServices(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.streamingLocators.list(resourceGroupName, accountName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listsStreamingLocators().catch(console.error);
