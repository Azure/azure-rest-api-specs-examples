const { SynapseManagementClient } = require("@azure/arm-synapse");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a Sql pool data masking rule.
 *
 * @summary Creates or updates a Sql pool data masking rule.
 * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/DataMaskingRuleCreateOrUpdateDefaultMax.json
 */
async function createOrUpdateDataMaskingRuleForDefaultMax() {
  const subscriptionId =
    process.env["SYNAPSE_SUBSCRIPTION_ID"] || "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = process.env["SYNAPSE_RESOURCE_GROUP"] || "sqlcrudtest-6852";
  const workspaceName = "sqlcrudtest-2080";
  const sqlPoolName = "sqlcrudtest-331";
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
  const client = new SynapseManagementClient(credential, subscriptionId);
  const result = await client.dataMaskingRules.createOrUpdate(
    resourceGroupName,
    workspaceName,
    sqlPoolName,
    dataMaskingRuleName,
    parameters
  );
  console.log(result);
}
