const { DnsResolverManagementClient } = require("@azure/arm-dnsresolver");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes an inbound endpoint for a DNS resolver. WARNING: This operation cannot be undone.
 *
 * @summary Deletes an inbound endpoint for a DNS resolver. WARNING: This operation cannot be undone.
 * x-ms-original-file: specification/dnsresolver/resource-manager/Microsoft.Network/stable/2022-07-01/examples/InboundEndpoint_Delete.json
 */
async function deleteInboundEndpointForDnsResolver() {
  const subscriptionId = "abdd4249-9f34-4cc6-8e42-c2e32110603e";
  const resourceGroupName = "sampleResourceGroup";
  const dnsResolverName = "sampleDnsResolver";
  const inboundEndpointName = "sampleInboundEndpoint";
  const credential = new DefaultAzureCredential();
  const client = new DnsResolverManagementClient(credential, subscriptionId);
  const result = await client.inboundEndpoints.beginDeleteAndWait(
    resourceGroupName,
    dnsResolverName,
    inboundEndpointName
  );
  console.log(result);
}

deleteInboundEndpointForDnsResolver().catch(console.error);
