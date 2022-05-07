Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-sql_9.0.1/sdk/sql/arm-sql/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a database service objective.
 *
 * @summary Gets a database service objective.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2014-04-01/examples/ServiceObjectiveGet.json
 */
async function getAServiceObjective() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "group1";
  const serverName = "sqlcrudtest";
  const serviceObjectiveName = "29dd7459-4a7c-4e56-be22-f0adda49440d";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.serviceObjectives.get(
    resourceGroupName,
    serverName,
    serviceObjectiveName
  );
  console.log(result);
}

getAServiceObjective().catch(console.error);
```
