const { DnsResolverManagementClient } = require("@azure/arm-dnsresolver");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a virtual network link to a DNS forwarding ruleset.
 *
 * @summary Creates or updates a virtual network link to a DNS forwarding ruleset.
 * x-ms-original-file: specification/dnsresolver/resource-manager/Microsoft.Network/preview/2020-04-01-preview/examples/VirtualNetworkLink_Put.json
 */
async function upsertVirtualNetworkLinkToADnsForwardingRuleset() {
  const subscriptionId = "abdd4249-9f34-4cc6-8e42-c2e32110603e";
  const resourceGroupName = "sampleResourceGroup";
  const dnsForwardingRulesetName = "sampleDnsForwardingRuleset";
  const virtualNetworkLinkName = "sampleVirtualNetworkLink";
  const parameters = {
    metadata: { additionalProp1: "value1" },
    virtualNetwork: {
      id: "/subscriptions/0403cfa9-9659-4f33-9f30-1f191c51d111/resourceGroups/sampleVnetResourceGroupName/providers/Microsoft.Network/virtualNetworks/sampleVirtualNetwork",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new DnsResolverManagementClient(credential, subscriptionId);
  const result = await client.virtualNetworkLinks.beginCreateOrUpdateAndWait(
    resourceGroupName,
    dnsForwardingRulesetName,
    virtualNetworkLinkName,
    parameters
  );
  console.log(result);
}

upsertVirtualNetworkLinkToADnsForwardingRuleset().catch(console.error);
