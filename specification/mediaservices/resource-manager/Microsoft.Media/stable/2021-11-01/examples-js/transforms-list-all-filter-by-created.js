const { AzureMediaServices } = require("@azure/arm-mediaservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the Transforms in the account.
 *
 * @summary Lists the Transforms in the account.
 * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/transforms-list-all-filter-by-created.json
 */
async function listsTheTransformsFilterByCreated() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "contosoresources";
  const accountName = "contosomedia";
  const filter =
    "properties/created gt 2021-11-01T00:00:00.0000000Z and properties/created le 2021-11-01T00:00:10.0000000Z";
  const orderby = "properties/created";
  const options = { filter, orderby };
  const credential = new DefaultAzureCredential();
  const client = new AzureMediaServices(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.transforms.list(resourceGroupName, accountName, options)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listsTheTransformsFilterByCreated().catch(console.error);
