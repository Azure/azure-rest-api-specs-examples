const { DnsResolverManagementClient } = require("@azure/arm-dnsresolver");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a virtual network link to a DNS forwarding ruleset. WARNING: This operation cannot be undone.
 *
 * @summary Deletes a virtual network link to a DNS forwarding ruleset. WARNING: This operation cannot be undone.
 * x-ms-original-file: specification/dnsresolver/resource-manager/Microsoft.Network/stable/2022-07-01/examples/VirtualNetworkLink_Delete.json
 */
async function deleteVirtualNetworkLinkToADnsForwardingRuleset() {
  const subscriptionId = "abdd4249-9f34-4cc6-8e42-c2e32110603e";
  const resourceGroupName = "sampleResourceGroup";
  const dnsForwardingRulesetName = "sampleDnsForwardingRuleset";
  const virtualNetworkLinkName = "sampleVirtualNetworkLink";
  const credential = new DefaultAzureCredential();
  const client = new DnsResolverManagementClient(credential, subscriptionId);
  const result = await client.virtualNetworkLinks.beginDeleteAndWait(
    resourceGroupName,
    dnsForwardingRulesetName,
    virtualNetworkLinkName
  );
  console.log(result);
}

deleteVirtualNetworkLinkToADnsForwardingRuleset().catch(console.error);
