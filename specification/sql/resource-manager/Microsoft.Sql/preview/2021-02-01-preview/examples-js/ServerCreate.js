const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a server.
 *
 * @summary Creates or updates a server.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2021-02-01-preview/examples/ServerCreate.json
 */
async function createServer() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "sqlcrudtest-7398";
  const serverName = "sqlcrudtest-4645";
  const parameters = {
    administratorLogin: "dummylogin",
    administratorLoginPassword: "PLACEHOLDER",
    administrators: {
      azureADOnlyAuthentication: true,
      login: "bob@contoso.com",
      principalType: "User",
      sid: "00000011-1111-2222-2222-123456789111",
      tenantId: "00000011-1111-2222-2222-123456789111",
    },
    location: "Japan East",
    publicNetworkAccess: "Enabled",
    restrictOutboundNetworkAccess: "Enabled",
  };
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.servers.beginCreateOrUpdateAndWait(
    resourceGroupName,
    serverName,
    parameters
  );
  console.log(result);
}

createServer().catch(console.error);
