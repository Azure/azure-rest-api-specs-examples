const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns a server communication link.
 *
 * @summary Returns a server communication link.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2014-04-01/examples/ServerCommunicationLinkGet.json
 */
async function getAServerCommunicationLink() {
  const subscriptionId =
    process.env["SQL_SUBSCRIPTION_ID"] || "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = process.env["SQL_RESOURCE_GROUP"] || "sqlcrudtest-7398";
  const serverName = "sqlcrudtest-4645";
  const communicationLinkName = "link1";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.serverCommunicationLinks.get(
    resourceGroupName,
    serverName,
    communicationLinkName,
  );
  console.log(result);
}
