const { IoTOperationsClient } = require("@azure/arm-iotoperations");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create a DataflowEndpointResource
 *
 * @summary create a DataflowEndpointResource
 * x-ms-original-file: 2024-09-15-preview/DataflowEndpoint_CreateOrUpdate_Kafka.json
 */
async function dataflowEndpointCreateOrUpdateKafka() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "F8C729F9-DF9C-4743-848F-96EE433D8E53";
  const client = new IoTOperationsClient(credential, subscriptionId);
  const result = await client.dataflowEndpoint.createOrUpdate(
    "rgiotoperations",
    "resource-name123",
    "generic-kafka-endpoint",
    {
      properties: {
        endpointType: "Kafka",
        kafkaSettings: {
          host: "example.kafka.local:9093",
          authentication: {
            method: "Sasl",
            saslSettings: { saslType: "Plain", secretRef: "my-secret" },
          },
          tls: {
            mode: "Enabled",
            trustedCaCertificateConfigMapRef: "ca-certificates",
          },
          consumerGroupId: "dataflows",
          compression: "Gzip",
          batching: {
            mode: "Enabled",
            latencyMs: 5,
            maxBytes: 1000000,
            maxMessages: 100000,
          },
          partitionStrategy: "Default",
          kafkaAcks: "All",
          copyMqttProperties: "Enabled",
          cloudEventAttributes: "Propagate",
        },
      },
      extendedLocation: {
        name: "qmbrfwcpwwhggszhrdjv",
        type: "CustomLocation",
      },
    },
  );
  console.log(result);
}
