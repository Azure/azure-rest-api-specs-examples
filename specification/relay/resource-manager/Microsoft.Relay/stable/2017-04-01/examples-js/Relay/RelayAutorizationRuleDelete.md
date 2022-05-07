Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-relay_3.0.1/sdk/relay/arm-relay/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { RelayAPI } = require("@azure/arm-relay");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a WCF relay authorization rule.
 *
 * @summary Deletes a WCF relay authorization rule.
 * x-ms-original-file: specification/relay/resource-manager/Microsoft.Relay/stable/2017-04-01/examples/Relay/RelayAutorizationRuleDelete.json
 */
async function relayAutorizationRuleDelete() {
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const resourceGroupName = "resourcegroup";
  const namespaceName = "example-RelayNamespace-01";
  const relayName = "example-Relay-wcf-01";
  const authorizationRuleName = "example-RelayAuthRules-01";
  const credential = new DefaultAzureCredential();
  const client = new RelayAPI(credential, subscriptionId);
  const result = await client.wCFRelays.deleteAuthorizationRule(
    resourceGroupName,
    namespaceName,
    relayName,
    authorizationRuleName
  );
  console.log(result);
}

relayAutorizationRuleDelete().catch(console.error);
```
