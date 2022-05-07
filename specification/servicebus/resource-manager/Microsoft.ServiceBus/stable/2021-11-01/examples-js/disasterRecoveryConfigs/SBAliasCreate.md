Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-servicebus_6.0.0/sdk/servicebus/arm-servicebus/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { ServiceBusManagementClient } = require("@azure/arm-servicebus");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a new Alias(Disaster Recovery configuration)
 *
 * @summary Creates or updates a new Alias(Disaster Recovery configuration)
 * x-ms-original-file: specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2021-11-01/examples/disasterRecoveryConfigs/SBAliasCreate.json
 */
async function sbAliasCreate() {
  const subscriptionId = "5f750a97-50d9-4e36-8081-c9ee4c0210d4";
  const resourceGroupName = "ardsouzatestRG";
  const namespaceName = "sdk-Namespace-8860";
  const alias = "sdk-Namespace-8860";
  const parameters = {
    alternateName: "alternameforAlias-Namespace-8860",
    partnerNamespace: "sdk-Namespace-37",
  };
  const credential = new DefaultAzureCredential();
  const client = new ServiceBusManagementClient(credential, subscriptionId);
  const result = await client.disasterRecoveryConfigs.createOrUpdate(
    resourceGroupName,
    namespaceName,
    alias,
    parameters
  );
  console.log(result);
}

sbAliasCreate().catch(console.error);
```
