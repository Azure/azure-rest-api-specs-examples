const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Gets the list of resources that are onboarded with NSP. These resources can be associated with a network security perimeter
 *
 * @summary Gets the list of resources that are onboarded with NSP. These resources can be associated with a network security perimeter
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/PerimeterAssociableResourcesList.json
 */
async function networkSecurityPerimeterAssociableResourceTypes() {
  const subscriptionId = process.env["NETWORK_SUBSCRIPTION_ID"] || "subId";
  const location = "westus";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.networkSecurityPerimeterAssociableResourceTypes.list(location)) {
    resArray.push(item);
  }
  console.log(resArray);
}
