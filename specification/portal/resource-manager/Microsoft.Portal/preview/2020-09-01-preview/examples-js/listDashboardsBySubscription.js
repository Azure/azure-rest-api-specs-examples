const { Portal } = require("@azure/arm-portal");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all the dashboards within a subscription.
 *
 * @summary Gets all the dashboards within a subscription.
 * x-ms-original-file: specification/portal/resource-manager/Microsoft.Portal/preview/2020-09-01-preview/examples/listDashboardsBySubscription.json
 */
async function listAllCustomResourceProvidersOnTheSubscription() {
  const subscriptionId =
    process.env["PORTAL_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new Portal(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.dashboards.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}
