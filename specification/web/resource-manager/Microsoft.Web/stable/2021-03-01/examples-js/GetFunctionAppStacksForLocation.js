const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Description for Get available Function app frameworks and their versions for location
 *
 * @summary Description for Get available Function app frameworks and their versions for location
 * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/GetFunctionAppStacksForLocation.json
 */
async function getLocationsFunctionAppStacks() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const location = "westus";
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.provider.listFunctionAppStacksForLocation(location)) {
    resArray.push(item);
  }
  console.log(resArray);
}

getLocationsFunctionAppStacks().catch(console.error);
