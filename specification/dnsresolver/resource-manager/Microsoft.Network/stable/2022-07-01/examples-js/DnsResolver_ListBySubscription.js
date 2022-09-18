const { DnsResolverManagementClient } = require("@azure/arm-dnsresolver");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists DNS resolvers in all resource groups of a subscription.
 *
 * @summary Lists DNS resolvers in all resource groups of a subscription.
 * x-ms-original-file: specification/dnsresolver/resource-manager/Microsoft.Network/stable/2022-07-01/examples/DnsResolver_ListBySubscription.json
 */
async function listDnsResolversBySubscription() {
  const subscriptionId = "abdd4249-9f34-4cc6-8e42-c2e32110603e";
  const credential = new DefaultAzureCredential();
  const client = new DnsResolverManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.dnsResolvers.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

listDnsResolversBySubscription().catch(console.error);
