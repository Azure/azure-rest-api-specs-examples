const { AzureMigrateV2 } = require("@azure/arm-migrate");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete the group from the project. The machines remain in the project. Deleting a non-existent group results in a no-operation.

A group is an aggregation mechanism for machines in a project. Therefore, deleting group does not delete machines in it.

 *
 * @summary Delete the group from the project. The machines remain in the project. Deleting a non-existent group results in a no-operation.

A group is an aggregation mechanism for machines in a project. Therefore, deleting group does not delete machines in it.

 * x-ms-original-file: specification/migrate/resource-manager/Microsoft.Migrate/stable/2019-10-01/examples/Groups_Delete.json
 */
async function groupsDelete() {
  const subscriptionId =
    process.env["MIGRATE_SUBSCRIPTION_ID"] || "6393a73f-8d55-47ef-b6dd-179b3e0c7910";
  const resourceGroupName = process.env["MIGRATE_RESOURCE_GROUP"] || "abgoyal-westEurope";
  const projectName = "abgoyalWEselfhostb72bproject";
  const groupName = "Test1";
  const credential = new DefaultAzureCredential();
  const client = new AzureMigrateV2(credential, subscriptionId);
  const result = await client.groups.delete(resourceGroupName, projectName, groupName);
  console.log(result);
}
