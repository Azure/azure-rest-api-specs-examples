Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-eventhub_5.0.1/sdk/eventhub/arm-eventhub/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { EventHubManagementClient } = require("@azure/arm-eventhub");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets an AuthorizationRule for a Namespace by rule name.
 *
 * @summary Gets an AuthorizationRule for a Namespace by rule name.
 * x-ms-original-file: specification/eventhub/resource-manager/Microsoft.EventHub/stable/2021-11-01/examples/disasterRecoveryConfigs/EHAliasAuthorizationRuleGet.json
 */
async function nameSpaceAuthorizationRuleGet() {
  const subscriptionId = "exampleSubscriptionId";
  const resourceGroupName = "exampleResourceGroup";
  const namespaceName = "sdk-Namespace-9080";
  const alias = "sdk-DisasterRecovery-4879";
  const authorizationRuleName = "sdk-Authrules-4879";
  const credential = new DefaultAzureCredential();
  const client = new EventHubManagementClient(credential, subscriptionId);
  const result = await client.disasterRecoveryConfigs.getAuthorizationRule(
    resourceGroupName,
    namespaceName,
    alias,
    authorizationRuleName
  );
  console.log(result);
}

nameSpaceAuthorizationRuleGet().catch(console.error);
```
