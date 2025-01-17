const { DeviceRegistryManagementClient } = require("@azure/arm-deviceregistry");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to delete a SchemaRegistry
 *
 * @summary delete a SchemaRegistry
 * x-ms-original-file: 2024-09-01-preview/Delete_SchemaRegistry.json
 */
async function deleteSchemaRegistry() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new DeviceRegistryManagementClient(credential, subscriptionId);
  await client.schemaRegistries.delete("myResourceGroup", "my-schema-registry");
}
