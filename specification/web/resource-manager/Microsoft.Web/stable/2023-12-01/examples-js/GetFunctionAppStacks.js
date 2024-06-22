const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Description for Get available Function app frameworks and their versions
 *
 * @summary Description for Get available Function app frameworks and their versions
 * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2023-12-01/examples/GetFunctionAppStacks.json
 */
async function getFunctionAppStacks() {
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential);
  const resArray = new Array();
  for await (let item of client.provider.listFunctionAppStacks()) {
    resArray.push(item);
  }
  console.log(resArray);
}
