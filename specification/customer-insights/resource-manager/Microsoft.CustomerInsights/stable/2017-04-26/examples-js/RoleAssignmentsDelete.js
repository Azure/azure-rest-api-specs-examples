const { CustomerInsightsManagementClient } = require("@azure/arm-customerinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the role assignment in the hub.
 *
 * @summary Deletes the role assignment in the hub.
 * x-ms-original-file: specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/RoleAssignmentsDelete.json
 */
async function roleAssignmentsDelete() {
  const subscriptionId = "subid";
  const resourceGroupName = "TestHubRG";
  const hubName = "sdkTestHub";
  const assignmentName = "assignmentName8976";
  const credential = new DefaultAzureCredential();
  const client = new CustomerInsightsManagementClient(credential, subscriptionId);
  const result = await client.roleAssignments.delete(resourceGroupName, hubName, assignmentName);
  console.log(result);
}

roleAssignmentsDelete().catch(console.error);
