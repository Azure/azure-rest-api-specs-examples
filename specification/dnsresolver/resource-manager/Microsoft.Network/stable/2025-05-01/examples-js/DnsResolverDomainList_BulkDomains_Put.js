const { DnsResolverManagementClient } = require("@azure/arm-dnsresolver");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Creates or updates a DNS resolver domain list.
 *
 * @summary Creates or updates a DNS resolver domain list.
 * x-ms-original-file: specification/dnsresolver/resource-manager/Microsoft.Network/stable/2025-05-01/examples/DnsResolverDomainList_BulkDomains_Put.json
 */
async function upsertDnsResolverDomainListWithBulkNumberOfDomains() {
  const subscriptionId =
    process.env["DNSRESOLVER_SUBSCRIPTION_ID"] || "abdd4249-9f34-4cc6-8e42-c2e32110603e";
  const resourceGroupName = process.env["DNSRESOLVER_RESOURCE_GROUP"] || "sampleResourceGroup";
  const dnsResolverDomainListName = "sampleDnsResolverDomainList";
  const parameters = {
    location: "westus2",
    tags: { key1: "value1" },
  };
  const credential = new DefaultAzureCredential();
  const client = new DnsResolverManagementClient(credential, subscriptionId);
  const result = await client.dnsResolverDomainLists.beginCreateOrUpdateAndWait(
    resourceGroupName,
    dnsResolverDomainListName,
    parameters,
  );
  console.log(result);
}
