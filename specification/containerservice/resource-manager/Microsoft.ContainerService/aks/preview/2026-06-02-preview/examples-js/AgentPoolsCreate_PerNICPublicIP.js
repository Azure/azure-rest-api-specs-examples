const { ContainerServiceClient } = require("@azure/arm-containerservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates an agent pool in the specified managed cluster.
 *
 * @summary creates or updates an agent pool in the specified managed cluster.
 * x-ms-original-file: 2026-06-02-preview/AgentPoolsCreate_PerNICPublicIP.json
 */
async function createAgentPoolWithPerNICPublicIPConfiguration() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ContainerServiceClient(credential, subscriptionId);
  const result = await client.agentPools.createOrUpdate("rg1", "clustername1", "agentpool1", {
    count: 3,
    orchestratorVersion: "",
    osType: "Linux",
    vmSize: "Standard_D8s_v3",
    networkProfile: {
      secondaryNetworkInterfaces: [
        {
          type: "Standard",
          vnetSubnetId:
            "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/myVNet/subnets/secondary-subnet-1",
          enableAcceleratedNetworking: true,
          publicIPAddressConfiguration: {
            publicIPAddressVersion: "IPv4",
            ipTags: [{ ipTagType: "FirstPartyUsage", tag: "teams" }],
          },
        },
        {
          type: "Standard",
          vnetSubnetId:
            "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/myVNet/subnets/secondary-subnet-2",
          enableAcceleratedNetworking: true,
          publicIPAddressConfiguration: {
            publicIPAddressVersion: "IPv4",
            publicIPPrefixID:
              "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/publicIPPrefixes/myPrefix",
          },
        },
      ],
    },
  });
  console.log(result);
}
