const { DnsResolverManagementClient } = require("@azure/arm-dnsresolver");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets properties of a DNS resolver.
 *
 * @summary Gets properties of a DNS resolver.
 * x-ms-original-file: specification/dnsresolver/resource-manager/Microsoft.Network/preview/2020-04-01-preview/examples/DnsResolver_Get.json
 */
async function retrieveDnsResolver() {
  const subscriptionId = "abdd4249-9f34-4cc6-8e42-c2e32110603e";
  const resourceGroupName = "sampleResourceGroup";
  const dnsResolverName = "sampleDnsResolver";
  const credential = new DefaultAzureCredential();
  const client = new DnsResolverManagementClient(credential, subscriptionId);
  const result = await client.dnsResolvers.get(resourceGroupName, dnsResolverName);
  console.log(result);
}

retrieveDnsResolver().catch(console.error);
