const { SynapseManagementClient } = require("@azure/arm-synapse");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get Sql Pool schema
 *
 * @summary Get Sql Pool schema
 * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/SqlPoolSchemaGet.json
 */
async function getDatabaseSchema() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "myRG";
  const workspaceName = "serverName";
  const sqlPoolName = "myDatabase";
  const schemaName = "dbo";
  const credential = new DefaultAzureCredential();
  const client = new SynapseManagementClient(credential, subscriptionId);
  const result = await client.sqlPoolSchemas.get(
    resourceGroupName,
    workspaceName,
    sqlPoolName,
    schemaName
  );
  console.log(result);
}

getDatabaseSchema().catch(console.error);
