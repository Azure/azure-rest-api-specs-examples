const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Description for Gets all available operations for the Microsoft.Web resource provider. Also exposes resource metric definitions
 *
 * @summary Description for Gets all available operations for the Microsoft.Web resource provider. Also exposes resource metric definitions
 * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2023-12-01/examples/ListOperations.json
 */
async function listOperations() {
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential);
  const resArray = new Array();
  for await (let item of client.provider.listOperations()) {
    resArray.push(item);
  }
  console.log(resArray);
}
