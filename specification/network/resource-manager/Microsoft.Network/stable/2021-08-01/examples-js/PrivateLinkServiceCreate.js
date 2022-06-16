const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates an private link service in the specified resource group.
 *
 * @summary Creates or updates an private link service in the specified resource group.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/PrivateLinkServiceCreate.json
 */
async function createPrivateLinkService() {
  const subscriptionId = "subId";
  const resourceGroupName = "rg1";
  const serviceName = "testPls";
  const parameters = {
    autoApproval: { subscriptions: ["subscription1", "subscription2"] },
    fqdns: ["fqdn1", "fqdn2", "fqdn3"],
    ipConfigurations: [
      {
        name: "fe-lb",
        privateIPAddress: "10.0.1.4",
        privateIPAddressVersion: "IPv4",
        privateIPAllocationMethod: "Static",
        subnet: {
          id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnetlb/subnets/subnetlb",
        },
      },
    ],
    loadBalancerFrontendIpConfigurations: [
      {
        id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/frontendIPConfigurations/fe-lb",
      },
    ],
    location: "eastus",
    visibility: {
      subscriptions: ["subscription1", "subscription2", "subscription3"],
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.privateLinkServices.beginCreateOrUpdateAndWait(
    resourceGroupName,
    serviceName,
    parameters
  );
  console.log(result);
}

createPrivateLinkService().catch(console.error);
