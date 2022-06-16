const { EventHubManagementClient } = require("@azure/arm-eventhub");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to
 *
 * @summary
 * x-ms-original-file: specification/eventhub/resource-manager/Microsoft.EventHub/stable/2021-11-01/examples/SchemaRegistry/SchemaRegistryCreate.json
 */
async function schemaRegistryCreate() {
  const subscriptionId = "e8baea74-64ce-459b-bee3-5aa4c47b3ae3";
  const resourceGroupName = "alitest";
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
    parameters
  );
  console.log(result);
}

schemaRegistryCreate().catch(console.error);
