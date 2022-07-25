const { DnsResolverManagementClient } = require("@azure/arm-dnsresolver");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists outbound endpoints for a DNS resolver.
 *
 * @summary Lists outbound endpoints for a DNS resolver.
 * x-ms-original-file: specification/dnsresolver/resource-manager/Microsoft.Network/preview/2020-04-01-preview/examples/OutboundEndpoint_List.json
 */
async function listOutboundEndpointsByDnsResolver() {
  const subscriptionId = "abdd4249-9f34-4cc6-8e42-c2e32110603e";
  const resourceGroupName = "sampleResourceGroup";
  const dnsResolverName = "sampleDnsResolver";
  const credential = new DefaultAzureCredential();
  const client = new DnsResolverManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.outboundEndpoints.list(resourceGroupName, dnsResolverName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listOutboundEndpointsByDnsResolver().catch(console.error);
