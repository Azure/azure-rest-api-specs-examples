Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-sql_9.0.1/sdk/sql/arm-sql/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a database data masking rule.
 *
 * @summary Creates or updates a database data masking rule.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2014-04-01/examples/DataMaskingRuleCreateOrUpdateDefaultMax.json
 */
async function createOrUpdateDataMaskingRuleForDefaultMax() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "sqlcrudtest-6852";
  const serverName = "sqlcrudtest-2080";
  const databaseName = "sqlcrudtest-331";
  const dataMaskingRuleName = "rule1";
  const parameters = {
    aliasName: "nickname",
    columnName: "test1",
    maskingFunction: "Default",
    ruleState: "Enabled",
    schemaName: "dbo",
    tableName: "Table_1",
  };
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.dataMaskingRules.createOrUpdate(
    resourceGroupName,
    serverName,
    databaseName,
    dataMaskingRuleName,
    parameters
  );
  console.log(result);
}

createOrUpdateDataMaskingRuleForDefaultMax().catch(console.error);
```
