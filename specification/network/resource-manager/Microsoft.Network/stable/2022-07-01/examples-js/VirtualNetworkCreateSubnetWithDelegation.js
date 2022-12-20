const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a virtual network in the specified resource group.
 *
 * @summary Creates or updates a virtual network in the specified resource group.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-07-01/examples/VirtualNetworkCreateSubnetWithDelegation.json
 */
async function createVirtualNetworkWithDelegatedSubnets() {
  const subscriptionId = "subId";
  const resourceGroupName = "rg1";
  const virtualNetworkName = "test-vnet";
  const parameters = {
    addressSpace: { addressPrefixes: ["10.0.0.0/16"] },
    location: "westcentralus",
    subnets: [
      {
        name: "test-1",
        addressPrefix: "10.0.0.0/24",
        delegations: [
          {
            name: "myDelegation",
            serviceName: "Microsoft.Sql/managedInstances",
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

createVirtualNetworkWithDelegatedSubnets().catch(console.error);
