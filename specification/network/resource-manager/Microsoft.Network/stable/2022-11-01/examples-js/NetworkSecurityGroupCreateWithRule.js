const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a network security group in the specified resource group.
 *
 * @summary Creates or updates a network security group in the specified resource group.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-11-01/examples/NetworkSecurityGroupCreateWithRule.json
 */
async function createNetworkSecurityGroupWithRule() {
  const subscriptionId = process.env["NETWORK_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["NETWORK_RESOURCE_GROUP"] || "rg1";
  const networkSecurityGroupName = "testnsg";
  const parameters = {
    location: "eastus",
    securityRules: [
      {
        name: "rule1",
        access: "Allow",
        destinationAddressPrefix: "*",
        destinationPortRange: "80",
        direction: "Inbound",
        priority: 130,
        sourceAddressPrefix: "*",
        sourcePortRange: "*",
        protocol: "*",
      },
    ],
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.networkSecurityGroups.beginCreateOrUpdateAndWait(
    resourceGroupName,
    networkSecurityGroupName,
    parameters
  );
  console.log(result);
}
