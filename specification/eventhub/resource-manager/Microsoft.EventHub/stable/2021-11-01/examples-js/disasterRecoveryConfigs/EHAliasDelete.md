Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-eventhub_5.0.1/sdk/eventhub/arm-eventhub/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { EventHubManagementClient } = require("@azure/arm-eventhub");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes an Alias(Disaster Recovery configuration)
 *
 * @summary Deletes an Alias(Disaster Recovery configuration)
 * x-ms-original-file: specification/eventhub/resource-manager/Microsoft.EventHub/stable/2021-11-01/examples/disasterRecoveryConfigs/EHAliasDelete.json
 */
async function ehAliasDelete() {
  const subscriptionId = "exampleSubscriptionId";
  const resourceGroupName = "exampleResourceGroup";
  const namespaceName = "sdk-Namespace-5849";
  const alias = "sdk-DisasterRecovery-3814";
  const credential = new DefaultAzureCredential();
  const client = new EventHubManagementClient(credential, subscriptionId);
  const result = await client.disasterRecoveryConfigs.delete(
    resourceGroupName,
    namespaceName,
    alias
  );
  console.log(result);
}

ehAliasDelete().catch(console.error);
```
