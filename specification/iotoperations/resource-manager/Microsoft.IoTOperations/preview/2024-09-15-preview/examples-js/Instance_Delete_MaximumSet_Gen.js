const { IoTOperationsClient } = require("@azure/arm-iotoperations");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to delete a InstanceResource
 *
 * @summary delete a InstanceResource
 * x-ms-original-file: 2024-09-15-preview/Instance_Delete_MaximumSet_Gen.json
 */
async function instanceDelete() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "F8C729F9-DF9C-4743-848F-96EE433D8E53";
  const client = new IoTOperationsClient(credential, subscriptionId);
  await client.instance.delete("rgiotoperations", "aio-instance");
}
