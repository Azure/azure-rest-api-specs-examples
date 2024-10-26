const { IoTOperationsClient } = require("@azure/arm-iotoperations");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create a DataflowEndpointResource
 *
 * @summary create a DataflowEndpointResource
 * x-ms-original-file: 2024-09-15-preview/DataflowEndpoint_CreateOrUpdate_AIO.json
 */
async function dataflowEndpointCreateOrUpdateAio() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "F8C729F9-DF9C-4743-848F-96EE433D8E53";
  const client = new IoTOperationsClient(credential, subscriptionId);
  const result = await client.dataflowEndpoint.createOrUpdate(
    "rgiotoperations",
    "resource-name123",
    "aio-builtin-broker-endpoint",
    {
      properties: {
        endpointType: "Mqtt",
        mqttSettings: {
          host: "aio-broker:18883",
          authentication: {
            method: "Kubernetes",
            serviceAccountTokenSettings: { audience: "aio-internal" },
          },
          tls: {
            mode: "Enabled",
            trustedCaCertificateConfigMapRef: "aio-ca-trust-bundle-test-only",
          },
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
