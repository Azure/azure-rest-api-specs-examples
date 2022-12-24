const { NetworkManagementClient } = require("@azure/arm-network-profile-2020-09-01-hybrid");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a virtual network in the specified resource group.
 *
 * @summary Creates or updates a virtual network in the specified resource group.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2018-11-01/examples/VirtualNetworkCreateSubnetWithDelegation.json
 */
async function createVirtualNetworkWithDelegatedSubnets() {
  const subscriptionId = process.env["NETWORK_SUBSCRIPTION_ID"] || "subId";
  const resourceGroupName = process.env["NETWORK_RESOURCE_GROUP"] || "rg1";
  const virtualNetworkName = "test-vnet";
  const parameters = {
    addressSpace: { addressPrefixes: ["10.0.0.0/16"] },
    subnets: [
      {
        name: "test-1",
        addressPrefix: "10.0.0.0/24",
        delegations: [
          {
            name: "myDelegation",
            serviceName: "Microsoft.Provider/resourceType",
          },
        ],
      },
    ],
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.virtualNetworks.beginCreateOrUpdateAndWait(
    resourceGroupName,
    virtualNetworkName,
    parameters
  );
  console.log(result);
}
