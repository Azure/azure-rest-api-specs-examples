```javascript
const { EventHubManagementClient } = require("@azure/arm-eventhub");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to
 *
 * @summary
 * x-ms-original-file: specification/eventhub/resource-manager/Microsoft.EventHub/stable/2021-11-01/examples/SchemaRegistry/SchemaRegistryGet.json
 */
async function schemaRegistryGet() {
  const subscriptionId = "e8baea74-64ce-459b-bee3-5aa4c47b3ae3";
  const resourceGroupName = "alitest";
  const namespaceName = "ali-ua-test-eh-system-1";
  const schemaGroupName = "testSchemaGroup1";
  const credential = new DefaultAzureCredential();
  const client = new EventHubManagementClient(credential, subscriptionId);
  const result = await client.schemaRegistry.get(resourceGroupName, namespaceName, schemaGroupName);
  console.log(result);
}

schemaRegistryGet().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-eventhub_5.0.1/sdk/eventhub/arm-eventhub/README.md) on how to add the SDK to your project and authenticate.
