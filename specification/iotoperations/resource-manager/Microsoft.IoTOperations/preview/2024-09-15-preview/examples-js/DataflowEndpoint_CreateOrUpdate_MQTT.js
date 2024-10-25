const { IoTOperationsClient } = require("@azure/arm-iotoperations");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create a DataflowEndpointResource
 *
 * @summary create a DataflowEndpointResource
 * x-ms-original-file: 2024-09-15-preview/DataflowEndpoint_CreateOrUpdate_MQTT.json
 */
async function dataflowEndpointCreateOrUpdateMqtt() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "F8C729F9-DF9C-4743-848F-96EE433D8E53";
  const client = new IoTOperationsClient(credential, subscriptionId);
  const result = await client.dataflowEndpoint.createOrUpdate(
    "rgiotoperations",
    "resource-name123",
    "generic-mqtt-broker-endpoint",
    {
      properties: {
        endpointType: "Mqtt",
        mqttSettings: {
          host: "example.broker.local:1883",
          authentication: {
            method: "X509Certificate",
            x509CertificateSettings: { secretRef: "example-secret" },
          },
          tls: { mode: "Disabled" },
          clientIdPrefix: "factory-gateway",
          retain: "Keep",
          sessionExpirySeconds: 3600,
          qos: 1,
          protocol: "WebSockets",
          maxInflightMessages: 100,
          keepAliveSeconds: 60,
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
