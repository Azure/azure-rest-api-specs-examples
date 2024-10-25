const { DnsResolverManagementClient } = require("@azure/arm-dnsresolver");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates a forwarding rule in a DNS forwarding ruleset.
 *
 * @summary Updates a forwarding rule in a DNS forwarding ruleset.
 * x-ms-original-file: specification/dnsresolver/resource-manager/Microsoft.Network/preview/2023-07-01-preview/examples/ForwardingRule_Patch.json
 */
async function updateForwardingRuleInADnsForwardingRuleset() {
  const subscriptionId =
    process.env["DNSRESOLVER_SUBSCRIPTION_ID"] || "abdd4249-9f34-4cc6-8e42-c2e32110603e";
  const resourceGroupName = process.env["DNSRESOLVER_RESOURCE_GROUP"] || "sampleResourceGroup";
  const dnsForwardingRulesetName = "sampleDnsForwardingRuleset";
  const forwardingRuleName = "sampleForwardingRule";
  const parameters = {
    forwardingRuleState: "Disabled",
    metadata: { additionalProp2: "value2" },
  };
  const credential = new DefaultAzureCredential();
  const client = new DnsResolverManagementClient(credential, subscriptionId);
  const result = await client.forwardingRules.update(
    resourceGroupName,
    dnsForwardingRulesetName,
    forwardingRuleName,
    parameters,
  );
  console.log(result);
}
