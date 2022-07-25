const { DnsResolverManagementClient } = require("@azure/arm-dnsresolver");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets properties of a virtual network link to a DNS forwarding ruleset.
 *
 * @summary Gets properties of a virtual network link to a DNS forwarding ruleset.
 * x-ms-original-file: specification/dnsresolver/resource-manager/Microsoft.Network/preview/2020-04-01-preview/examples/VirtualNetworkLink_Get.json
 */
async function retrieveVirtualNetworkLinkToADnsForwardingRuleset() {
  const subscriptionId = "abdd4249-9f34-4cc6-8e42-c2e32110603e";
  const resourceGroupName = "sampleResourceGroup";
  const dnsForwardingRulesetName = "sampleDnsForwardingRuleset";
  const virtualNetworkLinkName = "sampleVirtualNetworkLink";
  const credential = new DefaultAzureCredential();
  const client = new DnsResolverManagementClient(credential, subscriptionId);
  const result = await client.virtualNetworkLinks.get(
    resourceGroupName,
    dnsForwardingRulesetName,
    virtualNetworkLinkName
  );
  console.log(result);
}

retrieveVirtualNetworkLinkToADnsForwardingRuleset().catch(console.error);
