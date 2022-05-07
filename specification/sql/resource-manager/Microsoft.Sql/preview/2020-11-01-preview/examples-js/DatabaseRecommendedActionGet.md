Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-sql_9.0.1/sdk/sql/arm-sql/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a database recommended action.
 *
 * @summary Gets a database recommended action.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/DatabaseRecommendedActionGet.json
 */
async function getDatabaseRecommendedAction() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "workloadinsight-demos";
  const serverName = "misosisvr";
  const databaseName = "IndexAdvisor_test_3";
  const advisorName = "CreateIndex";
  const recommendedActionName = "IR_[CRM]_[DataPoints]_4821CD2F9510D98184BB";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.databaseRecommendedActions.get(
    resourceGroupName,
    serverName,
    databaseName,
    advisorName,
    recommendedActionName
  );
  console.log(result);
}

getDatabaseRecommendedAction().catch(console.error);
```
