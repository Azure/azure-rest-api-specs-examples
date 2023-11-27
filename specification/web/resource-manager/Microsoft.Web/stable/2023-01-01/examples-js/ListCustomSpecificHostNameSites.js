const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get custom hostnames under this subscription
 *
 * @summary Get custom hostnames under this subscription
 * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2023-01-01/examples/ListCustomSpecificHostNameSites.json
 */
async function getSpecificCustomHostnameUnderSubscription() {
  const subscriptionId =
    process.env["APPSERVICE_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const hostname = "www.example.com";
  const options = { hostname };
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.listCustomHostNameSites(options)) {
    resArray.push(item);
  }
  console.log(resArray);
}
