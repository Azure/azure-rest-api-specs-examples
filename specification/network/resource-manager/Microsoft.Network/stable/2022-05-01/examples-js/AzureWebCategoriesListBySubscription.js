const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all the Azure Web Categories in a subscription.
 *
 * @summary Gets all the Azure Web Categories in a subscription.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/AzureWebCategoriesListBySubscription.json
 */
async function listAllAzureWebCategoriesForAGivenSubscription() {
  const subscriptionId = "4de8428a-4a92-4cea-90ff-b47128b8cab8";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.webCategories.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}

listAllAzureWebCategoriesForAGivenSubscription().catch(console.error);
