const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the Container Apps in a given subscription.
 *
 * @summary Get the Container Apps in a given subscription.
 * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2022-03-01/examples/ListContainerAppsBySubscription.json
 */
async function listContainerAppsByResourceGroup() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.containerApps.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}

listContainerAppsByResourceGroup().catch(console.error);
