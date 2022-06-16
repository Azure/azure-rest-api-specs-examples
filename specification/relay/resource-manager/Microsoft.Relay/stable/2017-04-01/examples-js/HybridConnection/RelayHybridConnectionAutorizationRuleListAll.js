const { RelayAPI } = require("@azure/arm-relay");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Authorization rules for a hybrid connection.
 *
 * @summary Authorization rules for a hybrid connection.
 * x-ms-original-file: specification/relay/resource-manager/Microsoft.Relay/stable/2017-04-01/examples/HybridConnection/RelayHybridConnectionAutorizationRuleListAll.json
 */
async function relayHybridConnectionAutorizationRuleListAll() {
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const resourceGroupName = "resourcegroup";
  const namespaceName = "example-RelayNamespace-01";
  const hybridConnectionName = "example-Relay-Hybrid-01";
  const credential = new DefaultAzureCredential();
  const client = new RelayAPI(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.hybridConnections.listAuthorizationRules(
    resourceGroupName,
    namespaceName,
    hybridConnectionName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

relayHybridConnectionAutorizationRuleListAll().catch(console.error);
