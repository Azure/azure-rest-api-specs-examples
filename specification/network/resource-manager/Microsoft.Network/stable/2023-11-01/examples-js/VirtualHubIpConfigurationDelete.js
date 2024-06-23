const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a VirtualHubIpConfiguration.
 *
 * @summary Deletes a VirtualHubIpConfiguration.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2023-11-01/examples/VirtualHubIpConfigurationDelete.json
 */
async function virtualHubIPConfigurationDelete() {
  const subscriptionId = process.env["NETWORK_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["NETWORK_RESOURCE_GROUP"] || "rg1";
  const virtualHubName = "hub1";
  const ipConfigName = "ipconfig1";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.virtualHubIpConfiguration.beginDeleteAndWait(
    resourceGroupName,
    virtualHubName,
    ipConfigName,
  );
  console.log(result);
}
