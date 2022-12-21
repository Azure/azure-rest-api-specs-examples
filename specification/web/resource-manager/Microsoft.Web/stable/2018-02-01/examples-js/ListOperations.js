const { WebSiteManagementClient } = require("@azure/arm-appservice-profile-2020-09-01-hybrid");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all available operations for the Microsoft.Web resource provider. Also exposes resource metric definitions
 *
 * @summary Gets all available operations for the Microsoft.Web resource provider. Also exposes resource metric definitions
 * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2018-02-01/examples/ListOperations.json
 */
async function listOperations() {
  const subscriptionId =
    process.env["APPSERVICE_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.provider.listOperations()) {
    resArray.push(item);
  }
  console.log(resArray);
}
