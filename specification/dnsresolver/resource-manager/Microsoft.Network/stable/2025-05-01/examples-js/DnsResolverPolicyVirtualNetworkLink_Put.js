const { DnsResolverManagementClient } = require("@azure/arm-dnsresolver");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Creates or updates a DNS resolver policy virtual network link.
 *
 * @summary Creates or updates a DNS resolver policy virtual network link.
 * x-ms-original-file: specification/dnsresolver/resource-manager/Microsoft.Network/stable/2025-05-01/examples/DnsResolverPolicyVirtualNetworkLink_Put.json
 */
async function upsertDnsResolverPolicyVirtualNetworkLink() {
  const subscriptionId =
    process.env["DNSRESOLVER_SUBSCRIPTION_ID"] || "abdd4249-9f34-4cc6-8e42-c2e32110603e";
  const resourceGroupName = process.env["DNSRESOLVER_RESOURCE_GROUP"] || "sampleResourceGroup";
  const dnsResolverPolicyName = "sampleDnsResolverPolicy";
  const dnsResolverPolicyVirtualNetworkLinkName = "sampleVirtualNetworkLink";
  const parameters = {
    location: "westus2",
    tags: { key1: "value1" },
    virtualNetwork: {
      id: "/subscriptions/0403cfa9-9659-4f33-9f30-1f191c51d111/resourceGroups/sampleVnetResourceGroupName/providers/Microsoft.Network/virtualNetworks/sampleVirtualNetwork",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new DnsResolverManagementClient(credential, subscriptionId);
  const result = await client.dnsResolverPolicyVirtualNetworkLinks.beginCreateOrUpdateAndWait(
    resourceGroupName,
    dnsResolverPolicyName,
    dnsResolverPolicyVirtualNetworkLinkName,
    parameters,
  );
  console.log(result);
}
