const { DnsResolverManagementClient } = require("@azure/arm-dnsresolver");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists virtual network links to a DNS forwarding ruleset.
 *
 * @summary Lists virtual network links to a DNS forwarding ruleset.
 * x-ms-original-file: specification/dnsresolver/resource-manager/Microsoft.Network/preview/2020-04-01-preview/examples/VirtualNetworkLink_List.json
 */
async function listVirtualNetworkLinksToADnsForwardingRuleset() {
  const subscriptionId = "abdd4249-9f34-4cc6-8e42-c2e32110603e";
  const resourceGroupName = "sampleResourceGroup";
  const dnsForwardingRulesetName = "sampleDnsForwardingRuleset";
  const credential = new DefaultAzureCredential();
  const client = new DnsResolverManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.virtualNetworkLinks.list(
    resourceGroupName,
    dnsForwardingRulesetName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

listVirtualNetworkLinksToADnsForwardingRuleset().catch(console.error);
