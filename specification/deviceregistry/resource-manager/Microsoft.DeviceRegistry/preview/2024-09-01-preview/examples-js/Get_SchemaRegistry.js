const { DeviceRegistryManagementClient } = require("@azure/arm-deviceregistry");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to get a SchemaRegistry
 *
 * @summary get a SchemaRegistry
 * x-ms-original-file: 2024-09-01-preview/Get_SchemaRegistry.json
 */
async function getSchemaRegistry() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new DeviceRegistryManagementClient(credential, subscriptionId);
  const result = await client.schemaRegistries.get("myResourceGroup", "my-schema-registry");
  console.log(result);
}
