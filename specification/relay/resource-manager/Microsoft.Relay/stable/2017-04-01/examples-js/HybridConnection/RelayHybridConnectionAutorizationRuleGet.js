const { RelayAPI } = require("@azure/arm-relay");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Hybrid connection authorization rule for a hybrid connection by name.
 *
 * @summary Hybrid connection authorization rule for a hybrid connection by name.
 * x-ms-original-file: specification/relay/resource-manager/Microsoft.Relay/stable/2017-04-01/examples/HybridConnection/RelayHybridConnectionAutorizationRuleGet.json
 */
async function relayHybridConnectionAutorizationRuleGet() {
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const resourceGroupName = "resourcegroup";
  const namespaceName = "example-RelayNamespace-01";
  const hybridConnectionName = "example-Relay-Hybrid-01";
  const authorizationRuleName = "example-RelayAuthRules-01";
  const credential = new DefaultAzureCredential();
  const client = new RelayAPI(credential, subscriptionId);
  const result = await client.hybridConnections.getAuthorizationRule(
    resourceGroupName,
    namespaceName,
    hybridConnectionName,
    authorizationRuleName
  );
  console.log(result);
}

relayHybridConnectionAutorizationRuleGet().catch(console.error);
