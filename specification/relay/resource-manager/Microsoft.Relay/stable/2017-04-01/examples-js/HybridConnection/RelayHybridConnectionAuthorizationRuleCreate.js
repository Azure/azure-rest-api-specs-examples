const { RelayAPI } = require("@azure/arm-relay");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates an authorization rule for a hybrid connection.
 *
 * @summary Creates or updates an authorization rule for a hybrid connection.
 * x-ms-original-file: specification/relay/resource-manager/Microsoft.Relay/stable/2017-04-01/examples/HybridConnection/RelayHybridConnectionAuthorizationRuleCreate.json
 */
async function relayHybridConnectionAuthorizationRuleCreate() {
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const resourceGroupName = "resourcegroup";
  const namespaceName = "example-RelayNamespace-01";
  const hybridConnectionName = "example-Relay-Hybrid-01";
  const authorizationRuleName = "example-RelayAuthRules-01";
  const parameters = { rights: ["Listen", "Send"] };
  const credential = new DefaultAzureCredential();
  const client = new RelayAPI(credential, subscriptionId);
  const result = await client.hybridConnections.createOrUpdateAuthorizationRule(
    resourceGroupName,
    namespaceName,
    hybridConnectionName,
    authorizationRuleName,
    parameters
  );
  console.log(result);
}

relayHybridConnectionAuthorizationRuleCreate().catch(console.error);
