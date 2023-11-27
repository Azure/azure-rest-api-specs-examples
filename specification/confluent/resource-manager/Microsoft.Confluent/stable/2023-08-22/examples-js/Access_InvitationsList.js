const { ConfluentManagementClient } = require("@azure/arm-confluent");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Organization accounts invitation details
 *
 * @summary Organization accounts invitation details
 * x-ms-original-file: specification/confluent/resource-manager/Microsoft.Confluent/stable/2023-08-22/examples/Access_InvitationsList.json
 */
async function accessInvitationsList() {
  const subscriptionId =
    process.env["CONFLUENT_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["CONFLUENT_RESOURCE_GROUP"] || "myResourceGroup";
  const organizationName = "myOrganization";
  const body = {
    searchFilters: {
      pageSize: "10",
      pageToken: "asc4fts4ft",
      status: "INVITE_STATUS_SENT",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new ConfluentManagementClient(credential, subscriptionId);
  const result = await client.access.listInvitations(resourceGroupName, organizationName, body);
  console.log(result);
}
