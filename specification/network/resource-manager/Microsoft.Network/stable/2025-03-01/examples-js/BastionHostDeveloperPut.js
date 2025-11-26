const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Creates or updates the specified Bastion Host.
 *
 * @summary Creates or updates the specified Bastion Host.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/BastionHostDeveloperPut.json
 */
async function createDeveloperBastionHost() {
  const subscriptionId = process.env["NETWORK_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["NETWORK_RESOURCE_GROUP"] || "rg2";
  const bastionHostName = "bastionhostdeveloper";
  const parameters = {
    ipConfigurations: [],
    networkAcls: { ipRules: [{ addressPrefix: "1.1.1.1/16" }] },
    virtualNetwork: {
      id: "/subscriptions/subid/resourceGroups/rg2/providers/Microsoft.Network/virtualNetworks/vnet2",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.bastionHosts.beginCreateOrUpdateAndWait(
    resourceGroupName,
    bastionHostName,
    parameters,
  );
  console.log(result);
}
