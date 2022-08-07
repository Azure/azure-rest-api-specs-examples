const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a load balancer backend address pool.
 *
 * @summary Creates or updates a load balancer backend address pool.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/LBBackendAddressPoolWithBackendAddressesPut.json
 */
async function updateLoadBalancerBackendPoolWithBackendAddressesContainingVirtualNetworkAndIPAddress() {
  const subscriptionId = "subid";
  const resourceGroupName = "testrg";
  const loadBalancerName = "lb";
  const backendAddressPoolName = "backend";
  const parameters = {
    loadBalancerBackendAddresses: [
      {
        name: "address1",
        ipAddress: "10.0.0.4",
        virtualNetwork: {
          id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnetlb",
        },
      },
      {
        name: "address2",
        ipAddress: "10.0.0.5",
        virtualNetwork: {
          id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnetlb",
        },
      },
    ],
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.loadBalancerBackendAddressPools.beginCreateOrUpdateAndWait(
    resourceGroupName,
    loadBalancerName,
    backendAddressPoolName,
    parameters
  );
  console.log(result);
}

updateLoadBalancerBackendPoolWithBackendAddressesContainingVirtualNetworkAndIPAddress().catch(
  console.error
);
