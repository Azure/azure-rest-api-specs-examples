const { EventHubManagementClient } = require("@azure/arm-eventhub");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes an ApplicationGroup for a Namespace.
 *
 * @summary Deletes an ApplicationGroup for a Namespace.
 * x-ms-original-file: specification/eventhub/resource-manager/Microsoft.EventHub/stable/2024-01-01/examples/ApplicationGroup/ApplicationGroupDelete.json
 */
async function applicationGroupDelete() {
  const subscriptionId =
    process.env["EVENTHUB_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["EVENTHUB_RESOURCE_GROUP"] || "contosotest";
  const namespaceName = "contoso-ua-test-eh-system-1";
  const applicationGroupName = "appGroup1";
  const credential = new DefaultAzureCredential();
  const client = new EventHubManagementClient(credential, subscriptionId);
  const result = await client.applicationGroupOperations.delete(
    resourceGroupName,
    namespaceName,
    applicationGroupName,
  );
  console.log(result);
}
