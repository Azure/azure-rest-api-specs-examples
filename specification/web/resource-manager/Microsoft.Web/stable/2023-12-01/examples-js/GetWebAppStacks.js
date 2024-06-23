const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Description for Get available Web app frameworks and their versions
 *
 * @summary Description for Get available Web app frameworks and their versions
 * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2023-12-01/examples/GetWebAppStacks.json
 */
async function getWebAppStacks() {
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential);
  const resArray = new Array();
  for await (let item of client.provider.listWebAppStacks()) {
    resArray.push(item);
  }
  console.log(resArray);
}
