const { RelayAPI } = require("@azure/arm-relay");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Regenerates the primary or secondary connection strings to the hybrid connection.
 *
 * @summary Regenerates the primary or secondary connection strings to the hybrid connection.
 * x-ms-original-file: specification/relay/resource-manager/Microsoft.Relay/stable/2017-04-01/examples/HybridConnection/RelayHybridConnectionAuthorizationRuleRegenrateKey.json
 */
async function relayHybridConnectionAuthorizationRuleRegenrateKey() {
  const subscriptionId =
    process.env["RELAY_SUBSCRIPTION_ID"] || "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const resourceGroupName = process.env["RELAY_RESOURCE_GROUP"] || "resourcegroup";
  const namespaceName = "example-RelayNamespace-01";
  const hybridConnectionName = "example-Relay-Hybrid-01";
  const authorizationRuleName = "example-RelayAuthRules-01";
  const parameters = { keyType: "PrimaryKey" };
  const credential = new DefaultAzureCredential();
  const client = new RelayAPI(credential, subscriptionId);
  const result = await client.hybridConnections.regenerateKeys(
    resourceGroupName,
    namespaceName,
    hybridConnectionName,
    authorizationRuleName,
    parameters
  );
  console.log(result);
}
