```javascript
const { RelayAPI } = require("@azure/arm-relay");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates an authorization rule for a WCF relay.
 *
 * @summary Creates or updates an authorization rule for a WCF relay.
 * x-ms-original-file: specification/relay/resource-manager/Microsoft.Relay/stable/2017-04-01/examples/Relay/RelayAuthorizationRuleCreate.json
 */
async function relayAuthorizationRuleCreate() {
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const resourceGroupName = "resourcegroup";
  const namespaceName = "example-RelayNamespace-01";
  const relayName = "example-Relay-wcf-01";
  const authorizationRuleName = "example-RelayAuthRules-01";
  const parameters = { rights: ["Listen", "Send"] };
  const credential = new DefaultAzureCredential();
  const client = new RelayAPI(credential, subscriptionId);
  const result = await client.wCFRelays.createOrUpdateAuthorizationRule(
    resourceGroupName,
    namespaceName,
    relayName,
    authorizationRuleName,
    parameters
  );
  console.log(result);
}

relayAuthorizationRuleCreate().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-relay_3.0.1/sdk/relay/arm-relay/README.md) on how to add the SDK to your project and authenticate.
