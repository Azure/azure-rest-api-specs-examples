const { EventHubManagementClient } = require("@azure/arm-eventhub");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes an EventHub schema group.
 *
 * @summary Deletes an EventHub schema group.
 * x-ms-original-file: specification/eventhub/resource-manager/Microsoft.EventHub/stable/2024-01-01/examples/SchemaRegistry/SchemaRegistryDelete.json
 */
async function schemaRegistryDelete() {
  const subscriptionId =
    process.env["EVENTHUB_SUBSCRIPTION_ID"] || "e8baea74-64ce-459b-bee3-5aa4c47b3ae3";
  const resourceGroupName = process.env["EVENTHUB_RESOURCE_GROUP"] || "alitest";
  const namespaceName = "ali-ua-test-eh-system-1";
  const schemaGroupName = "testSchemaGroup1";
  const credential = new DefaultAzureCredential();
  const client = new EventHubManagementClient(credential, subscriptionId);
  const result = await client.schemaRegistry.delete(
    resourceGroupName,
    namespaceName,
    schemaGroupName,
  );
  console.log(result);
}
