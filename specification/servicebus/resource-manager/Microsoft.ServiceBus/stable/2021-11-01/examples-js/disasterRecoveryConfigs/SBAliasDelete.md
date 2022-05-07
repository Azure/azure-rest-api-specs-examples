Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-servicebus_6.0.0/sdk/servicebus/arm-servicebus/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { ServiceBusManagementClient } = require("@azure/arm-servicebus");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes an Alias(Disaster Recovery configuration)
 *
 * @summary Deletes an Alias(Disaster Recovery configuration)
 * x-ms-original-file: specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2021-11-01/examples/disasterRecoveryConfigs/SBAliasDelete.json
 */
async function sbAliasDelete() {
  const subscriptionId = "5f750a97-50d9-4e36-8081-c9ee4c0210d4";
  const resourceGroupName = "SouthCentralUS";
  const namespaceName = "sdk-Namespace-8860";
  const alias = "sdk-DisasterRecovery-3814";
  const credential = new DefaultAzureCredential();
  const client = new ServiceBusManagementClient(credential, subscriptionId);
  const result = await client.disasterRecoveryConfigs.delete(
    resourceGroupName,
    namespaceName,
    alias
  );
  console.log(result);
}

sbAliasDelete().catch(console.error);
```
