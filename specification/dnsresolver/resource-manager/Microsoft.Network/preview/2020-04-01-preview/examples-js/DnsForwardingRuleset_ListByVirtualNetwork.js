const { DnsResolverManagementClient } = require("@azure/arm-dnsresolver");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists DNS forwarding ruleset resource IDs attached to a virtual network.
 *
 * @summary Lists DNS forwarding ruleset resource IDs attached to a virtual network.
 * x-ms-original-file: specification/dnsresolver/resource-manager/Microsoft.Network/preview/2020-04-01-preview/examples/DnsForwardingRuleset_ListByVirtualNetwork.json
 */
async function listDnsForwardingRulesetsByVirtualNetwork() {
  const subscriptionId = "abdd4249-9f34-4cc6-8e42-c2e32110603e";
  const resourceGroupName = "sampleResourceGroup";
  const virtualNetworkName = "sampleVirtualNetwork";
  const credential = new DefaultAzureCredential();
  const client = new DnsResolverManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.dnsForwardingRulesets.listByVirtualNetwork(
    resourceGroupName,
    virtualNetworkName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

listDnsForwardingRulesetsByVirtualNetwork().catch(console.error);
