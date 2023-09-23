const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Determines whether a resource can be created with the specified name.
 *
 * @summary Determines whether a resource can be created with the specified name.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2023-02-01-preview/examples/CheckNameAvailabilityServerAvailable.json
 */
async function checkForAServerNameThatIsAvailable() {
  const subscriptionId =
    process.env["SQL_SUBSCRIPTION_ID"] || "00000000-1111-2222-3333-444444444444";
  const parameters = {
    name: "server1",
    type: "Microsoft.Sql/servers",
  };
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.servers.checkNameAvailability(parameters);
  console.log(result);
}
