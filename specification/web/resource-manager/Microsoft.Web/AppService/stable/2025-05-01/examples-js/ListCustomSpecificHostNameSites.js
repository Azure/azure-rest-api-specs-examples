const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to get custom hostnames under this subscription
 *
 * @summary get custom hostnames under this subscription
 * x-ms-original-file: 2025-05-01/ListCustomSpecificHostNameSites.json
 */
async function getSpecificCustomHostnameUnderSubscription() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.listCustomHostNameSites({ hostname: "www.example.com" })) {
    resArray.push(item);
  }

  console.log(resArray);
}
