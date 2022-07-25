const { SynapseManagementClient } = require("@azure/arm-synapse");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a Sql pool data masking policy
 *
 * @summary Creates or updates a Sql pool data masking policy
 * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/DataMaskingPolicyCreateOrUpdateMin.json
 */
async function createOrUpdateDataMaskingPolicyMin() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "sqlcrudtest-6852";
  const workspaceName = "sqlcrudtest-2080";
  const sqlPoolName = "sqlcrudtest-331";
  const parameters = { dataMaskingState: "Enabled" };
  const credential = new DefaultAzureCredential();
  const client = new SynapseManagementClient(credential, subscriptionId);
  const result = await client.dataMaskingPolicies.createOrUpdate(
    resourceGroupName,
    workspaceName,
    sqlPoolName,
    parameters
  );
  console.log(result);
}

createOrUpdateDataMaskingPolicyMin().catch(console.error);
