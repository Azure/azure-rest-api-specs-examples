const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates a database data masking rule.
 *
 * @summary creates or updates a database data masking rule.
 * x-ms-original-file: 2025-01-01/DataMaskingRuleCreateOrUpdateNumber.json
 */
async function createOrUpdateDataMaskingRuleForNumbers() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.dataMaskingRules.createOrUpdate(
    "sqlcrudtest-6852",
    "sqlcrudtest-2080",
    "sqlcrudtest-331",
    "rule1",
    {
      columnName: "test1",
      maskingFunction: "Number",
      numberFrom: "0",
      numberTo: "2",
      schemaName: "dbo",
      tableName: "Table_1",
    },
  );
  console.log(result);
}
