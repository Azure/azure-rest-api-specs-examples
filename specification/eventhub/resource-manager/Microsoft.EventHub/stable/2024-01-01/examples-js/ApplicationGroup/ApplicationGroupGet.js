const { EventHubManagementClient } = require("@azure/arm-eventhub");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets an ApplicationGroup for a Namespace.
 *
 * @summary Gets an ApplicationGroup for a Namespace.
 * x-ms-original-file: specification/eventhub/resource-manager/Microsoft.EventHub/stable/2024-01-01/examples/ApplicationGroup/ApplicationGroupGet.json
 */
async function applicationGroupGet() {
  const subscriptionId =
    process.env["EVENTHUB_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["EVENTHUB_RESOURCE_GROUP"] || "contosotest";
  const namespaceName = "contoso-ua-test-eh-system-1";
  const applicationGroupName = "appGroup1";
  const credential = new DefaultAzureCredential();
  const client = new EventHubManagementClient(credential, subscriptionId);
  const result = await client.applicationGroupOperations.get(
    resourceGroupName,
    namespaceName,
    applicationGroupName,
  );
  console.log(result);
}
