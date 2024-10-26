const { DnsResolverManagementClient } = require("@azure/arm-dnsresolver");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets properties of a DNS security rule for a DNS resolver policy.
 *
 * @summary Gets properties of a DNS security rule for a DNS resolver policy.
 * x-ms-original-file: specification/dnsresolver/resource-manager/Microsoft.Network/preview/2023-07-01-preview/examples/DnsSecurityRule_Get.json
 */
async function retrieveDnsSecurityRuleForDnsResolverPolicy() {
  const subscriptionId =
    process.env["DNSRESOLVER_SUBSCRIPTION_ID"] || "abdd4249-9f34-4cc6-8e42-c2e32110603e";
  const resourceGroupName = process.env["DNSRESOLVER_RESOURCE_GROUP"] || "sampleResourceGroup";
  const dnsResolverPolicyName = "sampleDnsResolverPolicy";
  const dnsSecurityRuleName = "sampleDnsSecurityRule";
  const credential = new DefaultAzureCredential();
  const client = new DnsResolverManagementClient(credential, subscriptionId);
  const result = await client.dnsSecurityRules.get(
    resourceGroupName,
    dnsResolverPolicyName,
    dnsSecurityRuleName,
  );
  console.log(result);
}
