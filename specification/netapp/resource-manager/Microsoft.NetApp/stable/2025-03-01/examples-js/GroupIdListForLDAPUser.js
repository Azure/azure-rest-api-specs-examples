const { NetAppManagementClient } = require("@azure/arm-netapp");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Returns the list of group Ids for a specific LDAP User
 *
 * @summary Returns the list of group Ids for a specific LDAP User
 * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2025-03-01/examples/GroupIdListForLDAPUser.json
 */
async function getGroupIdListForUser() {
  const subscriptionId =
    process.env["NETAPP_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["NETAPP_RESOURCE_GROUP"] || "myRG";
  const accountName = "account1";
  const poolName = "pool1";
  const volumeName = "volume1";
  const body = { username: "user1" };
  const credential = new DefaultAzureCredential();
  const client = new NetAppManagementClient(credential, subscriptionId);
  const result = await client.volumes.beginListGetGroupIdListForLdapUserAndWait(
    resourceGroupName,
    accountName,
    poolName,
    volumeName,
    body,
  );
  console.log(result);
}
