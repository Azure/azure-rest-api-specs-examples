const { IoTOperationsClient } = require("@azure/arm-iotoperations");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create a DataflowProfileResource
 *
 * @summary create a DataflowProfileResource
 * x-ms-original-file: 2024-09-15-preview/DataflowProfile_CreateOrUpdate_MaximumSet_Gen.json
 */
async function dataflowProfileCreateOrUpdate() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "F8C729F9-DF9C-4743-848F-96EE433D8E53";
  const client = new IoTOperationsClient(credential, subscriptionId);
  const result = await client.dataflowProfile.createOrUpdate(
    "rgiotoperations",
    "resource-name123",
    "resource-name123",
    {
      properties: {
        diagnostics: {
          logs: { level: "rnmwokumdmebpmfxxxzvvjfdywotav" },
          metrics: { prometheusPort: 7581 },
        },
        instanceCount: 14,
      },
      extendedLocation: {
        name: "qmbrfwcpwwhggszhrdjv",
        type: "CustomLocation",
      },
    },
  );
  console.log(result);
}
