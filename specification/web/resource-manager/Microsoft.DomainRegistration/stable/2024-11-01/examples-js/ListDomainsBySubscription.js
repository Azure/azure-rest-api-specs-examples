const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Description for Get all domains in a subscription.
 *
 * @summary Description for Get all domains in a subscription.
 * x-ms-original-file: specification/web/resource-manager/Microsoft.DomainRegistration/stable/2024-11-01/examples/ListDomainsBySubscription.json
 */
async function listDomainsBySubscription() {
  const subscriptionId =
    process.env["APPSERVICE_SUBSCRIPTION_ID"] || "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.domains.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}
