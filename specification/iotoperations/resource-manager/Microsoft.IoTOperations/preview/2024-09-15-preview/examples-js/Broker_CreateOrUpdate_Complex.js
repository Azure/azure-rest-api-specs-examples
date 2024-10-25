const { IoTOperationsClient } = require("@azure/arm-iotoperations");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create a BrokerResource
 *
 * @summary create a BrokerResource
 * x-ms-original-file: 2024-09-15-preview/Broker_CreateOrUpdate_Complex.json
 */
async function brokerCreateOrUpdateComplex() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "F8C729F9-DF9C-4743-848F-96EE433D8E53";
  const client = new IoTOperationsClient(credential, subscriptionId);
  const result = await client.broker.createOrUpdate(
    "rgiotoperations",
    "resource-name123",
    "resource-name123",
    {
      properties: {
        cardinality: {
          backendChain: { partitions: 2, redundancyFactor: 2, workers: 2 },
          frontend: { replicas: 2, workers: 2 },
        },
        diskBackedMessageBuffer: { maxSize: "50M" },
        generateResourceLimits: { cpu: "Enabled" },
        memoryProfile: "Medium",
      },
      extendedLocation: {
        name: "qmbrfwcpwwhggszhrdjv",
        type: "CustomLocation",
      },
    },
  );
  console.log(result);
}
