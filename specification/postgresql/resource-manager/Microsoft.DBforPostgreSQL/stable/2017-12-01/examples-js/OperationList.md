Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-postgresql_6.0.1/sdk/postgresql/arm-postgresql/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { PostgreSQLManagementClient } = require("@azure/arm-postgresql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all of the available REST API operations.
 *
 * @summary Lists all of the available REST API operations.
 * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2017-12-01/examples/OperationList.json
 */
async function operationList() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new PostgreSQLManagementClient(credential, subscriptionId);
  const result = await client.operations.list();
  console.log(result);
}

operationList().catch(console.error);
```
