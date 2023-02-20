const { DataProtectionClient } = require("@azure/arm-dataprotection");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns ResourceGuards collection belonging to a subscription.
 *
 * @summary Returns ResourceGuards collection belonging to a subscription.
 * x-ms-original-file: specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2023-01-01/examples/ResourceGuardCRUD/GetResourceGuardsInSubscription.json
 */
async function getResourceGuardsInSubscription() {
  const subscriptionId =
    process.env["DATAPROTECTION_SUBSCRIPTION_ID"] || "0b352192-dcac-4cc7-992e-a96190ccc68c";
  const credential = new DefaultAzureCredential();
  const client = new DataProtectionClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.resourceGuards.listResourcesInSubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}
