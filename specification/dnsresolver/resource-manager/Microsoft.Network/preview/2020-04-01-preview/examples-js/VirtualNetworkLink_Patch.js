const { DnsResolverManagementClient } = require("@azure/arm-dnsresolver");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates a virtual network link to a DNS forwarding ruleset.
 *
 * @summary Updates a virtual network link to a DNS forwarding ruleset.
 * x-ms-original-file: specification/dnsresolver/resource-manager/Microsoft.Network/preview/2020-04-01-preview/examples/VirtualNetworkLink_Patch.json
 */
async function updateVirtualNetworkLinkToADnsForwardingRuleset() {
  const subscriptionId = "abdd4249-9f34-4cc6-8e42-c2e32110603e";
  const resourceGroupName = "sampleResourceGroup";
  const dnsForwardingRulesetName = "sampleDnsForwardingRuleset";
  const virtualNetworkLinkName = "sampleVirtualNetworkLink";
  const parameters = {
    metadata: { additionalProp1: "value1" },
  };
  const credential = new DefaultAzureCredential();
  const client = new DnsResolverManagementClient(credential, subscriptionId);
  const result = await client.virtualNetworkLinks.beginUpdateAndWait(
    resourceGroupName,
    dnsForwardingRulesetName,
    virtualNetworkLinkName,
    parameters
  );
  console.log(result);
}

updateVirtualNetworkLinkToADnsForwardingRuleset().catch(console.error);
