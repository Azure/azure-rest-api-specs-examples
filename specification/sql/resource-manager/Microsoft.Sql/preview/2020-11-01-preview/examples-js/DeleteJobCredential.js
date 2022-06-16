const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a job credential.
 *
 * @summary Deletes a job credential.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/DeleteJobCredential.json
 */
async function deleteACredential() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "group1";
  const serverName = "server1";
  const jobAgentName = "agent1";
  const credentialName = "cred1";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.jobCredentials.delete(
    resourceGroupName,
    serverName,
    jobAgentName,
    credentialName
  );
  console.log(result);
}

deleteACredential().catch(console.error);
