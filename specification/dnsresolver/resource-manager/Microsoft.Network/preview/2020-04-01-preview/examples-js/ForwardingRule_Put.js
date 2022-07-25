const { DnsResolverManagementClient } = require("@azure/arm-dnsresolver");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a forwarding rule in a DNS forwarding ruleset.
 *
 * @summary Creates or updates a forwarding rule in a DNS forwarding ruleset.
 * x-ms-original-file: specification/dnsresolver/resource-manager/Microsoft.Network/preview/2020-04-01-preview/examples/ForwardingRule_Put.json
 */
async function upsertForwardingRuleInADnsForwardingRuleset() {
  const subscriptionId = "abdd4249-9f34-4cc6-8e42-c2e32110603e";
  const resourceGroupName = "sampleResourceGroup";
  const dnsForwardingRulesetName = "sampleDnsForwardingRuleset";
  const forwardingRuleName = "sampleForwardingRule";
  const parameters = {
    domainName: "contoso.com.",
    forwardingRuleState: "Enabled",
    metadata: { additionalProp1: "value1" },
    targetDnsServers: [
      { ipAddress: "10.0.0.1", port: 53 },
      { ipAddress: "10.0.0.2", port: 53 },
    ],
  };
  const credential = new DefaultAzureCredential();
  const client = new DnsResolverManagementClient(credential, subscriptionId);
  const result = await client.forwardingRules.createOrUpdate(
    resourceGroupName,
    dnsForwardingRulesetName,
    forwardingRuleName,
    parameters
  );
  console.log(result);
}

upsertForwardingRuleInADnsForwardingRuleset().catch(console.error);
