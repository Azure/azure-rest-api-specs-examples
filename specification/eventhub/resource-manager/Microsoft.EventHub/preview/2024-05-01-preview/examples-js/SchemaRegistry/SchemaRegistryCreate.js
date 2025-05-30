const { EventHubManagementClient } = require("@azure/arm-eventhub");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Creates or Updates an EventHub schema group.
 *
 * @summary Creates or Updates an EventHub schema group.
 * x-ms-original-file: specification/eventhub/resource-manager/Microsoft.EventHub/preview/2024-05-01-preview/examples/SchemaRegistry/SchemaRegistryCreate.json
 */
async function schemaRegistryCreate() {
  const subscriptionId =
    process.env["EVENTHUB_SUBSCRIPTION_ID"] || "e8baea74-64ce-459b-bee3-5aa4c47b3ae3";
  const resourceGroupName = process.env["EVENTHUB_RESOURCE_GROUP"] || "alitest";
  const namespaceName = "ali-ua-test-eh-system-1";
  const schemaGroupName = "testSchemaGroup1";
  const parameters = {
    groupProperties: {},
    schemaCompatibility: "Forward",
    schemaType: "Avro",
  };
  const credential = new DefaultAzureCredential();
  const client = new EventHubManagementClient(credential, subscriptionId);
  const result = await client.schemaRegistry.createOrUpdate(
    resourceGroupName,
    namespaceName,
    schemaGroupName,
    parameters,
  );
  console.log(result);
}
