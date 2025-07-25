const { DnsResolverManagementClient } = require("@azure/arm-dnsresolver");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Lists DNS resolvers within a resource group.
 *
 * @summary Lists DNS resolvers within a resource group.
 * x-ms-original-file: specification/dnsresolver/resource-manager/Microsoft.Network/stable/2025-05-01/examples/DnsResolver_ListByResourceGroup.json
 */
async function listDnsResolversByResourceGroup() {
  const subscriptionId =
    process.env["DNSRESOLVER_SUBSCRIPTION_ID"] || "abdd4249-9f34-4cc6-8e42-c2e32110603e";
  const resourceGroupName = process.env["DNSRESOLVER_RESOURCE_GROUP"] || "sampleResourceGroup";
  const credential = new DefaultAzureCredential();
  const client = new DnsResolverManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.dnsResolvers.listByResourceGroup(resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
