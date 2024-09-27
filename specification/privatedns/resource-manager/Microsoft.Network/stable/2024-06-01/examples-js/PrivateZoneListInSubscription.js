const { PrivateDnsManagementClient } = require("@azure/arm-privatedns");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the Private DNS zones in all resource groups in a subscription.
 *
 * @summary Lists the Private DNS zones in all resource groups in a subscription.
 * x-ms-original-file: specification/privatedns/resource-manager/Microsoft.Network/stable/2024-06-01/examples/PrivateZoneListInSubscription.json
 */
async function getPrivateDnsZoneBySubscription() {
  const subscriptionId = process.env["PRIVATEDNS_SUBSCRIPTION_ID"] || "subscriptionId";
  const credential = new DefaultAzureCredential();
  const client = new PrivateDnsManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.privateZones.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}
