const { SynapseManagementClient } = require("@azure/arm-synapse");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a SQL pool's blob auditing policy.
 *
 * @summary Creates or updates a SQL pool's blob auditing policy.
 * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/CreateOrUpdateSqlPoolBlobAuditingWithMinParameters.json
 */
async function createOrUpdateADatabaseBlobAuditingPolicyWithMinimalParameters() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "blobauditingtest-4799";
  const workspaceName = "blobauditingtest-6440";
  const sqlPoolName = "testdb";
  const parameters = {
    state: "Enabled",
    storageAccountAccessKey:
      "sdlfkjabc+sdlfkjsdlkfsjdfLDKFTERLKFDFKLjsdfksjdflsdkfD2342309432849328476458/3RSD==",
    storageEndpoint: "https://mystorage.blob.core.windows.net",
  };
  const credential = new DefaultAzureCredential();
  const client = new SynapseManagementClient(credential, subscriptionId);
  const result = await client.sqlPoolBlobAuditingPolicies.createOrUpdate(
    resourceGroupName,
    workspaceName,
    sqlPoolName,
    parameters
  );
  console.log(result);
}

createOrUpdateADatabaseBlobAuditingPolicyWithMinimalParameters().catch(console.error);
