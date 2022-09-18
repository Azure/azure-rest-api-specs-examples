const { DnsResolverManagementClient } = require("@azure/arm-dnsresolver");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates an outbound endpoint for a DNS resolver.
 *
 * @summary Updates an outbound endpoint for a DNS resolver.
 * x-ms-original-file: specification/dnsresolver/resource-manager/Microsoft.Network/stable/2022-07-01/examples/OutboundEndpoint_Patch.json
 */
async function updateOutboundEndpointForDnsResolver() {
  const subscriptionId = "abdd4249-9f34-4cc6-8e42-c2e32110603e";
  const resourceGroupName = "sampleResourceGroup";
  const dnsResolverName = "sampleDnsResolver";
  const outboundEndpointName = "sampleOutboundEndpoint";
  const parameters = { tags: { key1: "value1" } };
  const credential = new DefaultAzureCredential();
  const client = new DnsResolverManagementClient(credential, subscriptionId);
  const result = await client.outboundEndpoints.beginUpdateAndWait(
    resourceGroupName,
    dnsResolverName,
    outboundEndpointName,
    parameters
  );
  console.log(result);
}

updateOutboundEndpointForDnsResolver().catch(console.error);
