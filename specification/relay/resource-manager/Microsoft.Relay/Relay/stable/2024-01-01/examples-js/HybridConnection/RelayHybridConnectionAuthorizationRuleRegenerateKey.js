const { RelayAPI } = require("@azure/arm-relay");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to regenerates the primary or secondary connection strings to the hybrid connection.
 *
 * @summary regenerates the primary or secondary connection strings to the hybrid connection.
 * x-ms-original-file: 2024-01-01/HybridConnection/RelayHybridConnectionAuthorizationRuleRegenerateKey.json
 */
async function relayHybridConnectionAuthorizationRuleRegenerateKey() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const client = new RelayAPI(credential, subscriptionId);
  const result = await client.hybridConnections.regenerateKeys(
    "resourcegroup",
    "example-RelayNamespace-01",
    "example-Relay-Hybrid-01",
    "example-RelayAuthRules-01",
    { keyType: "PrimaryKey" },
  );
  console.log(result);
}
