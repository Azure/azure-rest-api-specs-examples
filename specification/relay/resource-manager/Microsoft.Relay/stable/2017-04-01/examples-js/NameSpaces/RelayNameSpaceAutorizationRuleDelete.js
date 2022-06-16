const { RelayAPI } = require("@azure/arm-relay");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a namespace authorization rule.
 *
 * @summary Deletes a namespace authorization rule.
 * x-ms-original-file: specification/relay/resource-manager/Microsoft.Relay/stable/2017-04-01/examples/NameSpaces/RelayNameSpaceAutorizationRuleDelete.json
 */
async function relayNameSpaceAutorizationRuleDelete() {
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const resourceGroupName = "resourcegroup";
  const namespaceName = "example-RelayNamespace-01";
  const authorizationRuleName = "example-RelayAuthRules-01";
  const credential = new DefaultAzureCredential();
  const client = new RelayAPI(credential, subscriptionId);
  const result = await client.namespaces.deleteAuthorizationRule(
    resourceGroupName,
    namespaceName,
    authorizationRuleName
  );
  console.log(result);
}

relayNameSpaceAutorizationRuleDelete().catch(console.error);
