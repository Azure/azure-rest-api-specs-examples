const { DnsResolverManagementClient } = require("@azure/arm-dnsresolver");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates a DNS security rule.
 *
 * @summary Updates a DNS security rule.
 * x-ms-original-file: specification/dnsresolver/resource-manager/Microsoft.Network/preview/2023-07-01-preview/examples/DnsSecurityRule_Patch.json
 */
async function updateDnsSecurityRuleForDnsResolverPolicy() {
  const subscriptionId =
    process.env["DNSRESOLVER_SUBSCRIPTION_ID"] || "abdd4249-9f34-4cc6-8e42-c2e32110603e";
  const resourceGroupName = process.env["DNSRESOLVER_RESOURCE_GROUP"] || "sampleResourceGroup";
  const dnsResolverPolicyName = "sampleDnsResolverPolicy";
  const dnsSecurityRuleName = "sampleDnsSecurityRule";
  const parameters = {
    dnsSecurityRuleState: "Disabled",
    tags: { key1: "value1" },
  };
  const credential = new DefaultAzureCredential();
  const client = new DnsResolverManagementClient(credential, subscriptionId);
  const result = await client.dnsSecurityRules.beginUpdateAndWait(
    resourceGroupName,
    dnsResolverPolicyName,
    dnsSecurityRuleName,
    parameters,
  );
  console.log(result);
}
