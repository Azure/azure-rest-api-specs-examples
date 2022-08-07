const { AzureMediaServices } = require("@azure/arm-mediaservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List Media Services accounts in the resource group
 *
 * @summary List Media Services accounts in the resource group
 * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/accounts-list-all-accounts.json
 */
async function listAllMediaServicesAccounts() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "contoso";
  const credential = new DefaultAzureCredential();
  const client = new AzureMediaServices(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.mediaservices.list(resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listAllMediaServicesAccounts().catch(console.error);
