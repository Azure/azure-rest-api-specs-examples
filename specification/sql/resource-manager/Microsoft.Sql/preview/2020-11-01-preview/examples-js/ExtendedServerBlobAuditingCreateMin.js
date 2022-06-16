const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates an extended server's blob auditing policy.
 *
 * @summary Creates or updates an extended server's blob auditing policy.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ExtendedServerBlobAuditingCreateMin.json
 */
async function updateAServerExtendedBlobAuditingPolicyWithMinimalParameters() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "blobauditingtest-4799";
  const serverName = "blobauditingtest-6440";
  const parameters = {
    state: "Enabled",
    storageAccountAccessKey:
      "sdlfkjabc+sdlfkjsdlkfsjdfLDKFTERLKFDFKLjsdfksjdflsdkfD2342309432849328476458/3RSD==",
    storageEndpoint: "https://mystorage.blob.core.windows.net",
  };
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.extendedServerBlobAuditingPolicies.beginCreateOrUpdateAndWait(
    resourceGroupName,
    serverName,
    parameters
  );
  console.log(result);
}

updateAServerExtendedBlobAuditingPolicyWithMinimalParameters().catch(console.error);
