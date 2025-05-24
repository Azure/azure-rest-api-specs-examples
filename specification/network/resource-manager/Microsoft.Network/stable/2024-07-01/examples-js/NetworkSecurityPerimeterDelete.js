const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Deletes a network security perimeter.
 *
 * @summary Deletes a network security perimeter.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/NetworkSecurityPerimeterDelete.json
 */
async function networkSecurityPerimeterDelete() {
  const subscriptionId = process.env["NETWORK_SUBSCRIPTION_ID"] || "subId";
  const resourceGroupName = process.env["NETWORK_RESOURCE_GROUP"] || "rg1";
  const networkSecurityPerimeterName = "testNSP1";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.networkSecurityPerimeters.beginDeleteAndWait(
    resourceGroupName,
    networkSecurityPerimeterName,
  );
  console.log(result);
}
