const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to retrieves the available specified Network Virtual Appliance Inbound Security Rules Collection.
 *
 * @summary retrieves the available specified Network Virtual Appliance Inbound Security Rules Collection.
 * x-ms-original-file: 2026-05-01/InboundSecurityRuleSourceIPProtocolGet.json
 */
async function getAPermanentInternetIngressRuleWithSourceIPProtocolLoadDistribution() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.inboundSecurityRuleOperations.get("rg1", "nva", "rule1");
  console.log(result);
}
