const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Description for Get available Function app frameworks and their versions for location
 *
 * @summary Description for Get available Function app frameworks and their versions for location
 * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2024-11-01/examples/GetFunctionAppStacksForLocation.json
 */
async function getLocationsFunctionAppStacks() {
  const location = "westus";
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential);
  const resArray = new Array();
  for await (const item of client.provider.listFunctionAppStacksForLocation(location)) {
    resArray.push(item);
  }
  console.log(resArray);
}
