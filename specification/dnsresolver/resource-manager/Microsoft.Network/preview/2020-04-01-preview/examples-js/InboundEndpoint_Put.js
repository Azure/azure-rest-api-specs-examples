const { DnsResolverManagementClient } = require("@azure/arm-dnsresolver");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates an inbound endpoint for a DNS resolver.
 *
 * @summary Creates or updates an inbound endpoint for a DNS resolver.
 * x-ms-original-file: specification/dnsresolver/resource-manager/Microsoft.Network/preview/2020-04-01-preview/examples/InboundEndpoint_Put.json
 */
async function upsertInboundEndpointForDnsResolver() {
  const subscriptionId = "abdd4249-9f34-4cc6-8e42-c2e32110603e";
  const resourceGroupName = "sampleResourceGroup";
  const dnsResolverName = "sampleDnsResolver";
  const inboundEndpointName = "sampleInboundEndpoint";
  const parameters = {
    ipConfigurations: [
      {
        privateIpAddress: "255.255.255.255",
        privateIpAllocationMethod: "Static",
        subnet: {
          id: "/subscriptions/0403cfa9-9659-4f33-9f30-1f191c51d111/resourceGroups/sampleVnetResourceGroupName/providers/Microsoft.Network/virtualNetworks/sampleVirtualNetwork/subnets/sampleSubnet",
        },
      },
    ],
    location: "westus2",
    tags: { key1: "value1" },
  };
  const credential = new DefaultAzureCredential();
  const client = new DnsResolverManagementClient(credential, subscriptionId);
  const result = await client.inboundEndpoints.beginCreateOrUpdateAndWait(
    resourceGroupName,
    dnsResolverName,
    inboundEndpointName,
    parameters
  );
  console.log(result);
}

upsertInboundEndpointForDnsResolver().catch(console.error);
