Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-sql_9.0.1/sdk/sql/arm-sql/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Sets a database's long term retention policy.
 *
 * @summary Sets a database's long term retention policy.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/LongTermRetentionPolicyCreateOrUpdate.json
 */
async function createOrUpdateTheLongTermRetentionPolicyForTheDatabase() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "resourceGroup";
  const serverName = "testserver";
  const databaseName = "testDatabase";
  const policyName = "default";
  const parameters = {
    monthlyRetention: "P1Y",
    weekOfYear: 5,
    weeklyRetention: "P1M",
    yearlyRetention: "P5Y",
  };
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.longTermRetentionPolicies.beginCreateOrUpdateAndWait(
    resourceGroupName,
    serverName,
    databaseName,
    policyName,
    parameters
  );
  console.log(result);
}

createOrUpdateTheLongTermRetentionPolicyForTheDatabase().catch(console.error);
```
