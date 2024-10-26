const { IoTOperationsClient } = require("@azure/arm-iotoperations");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create a BrokerListenerResource
 *
 * @summary create a BrokerListenerResource
 * x-ms-original-file: 2024-09-15-preview/BrokerListener_CreateOrUpdate_Simple.json
 */
async function brokerListenerCreateOrUpdateSimple() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "F8C729F9-DF9C-4743-848F-96EE433D8E53";
  const client = new IoTOperationsClient(credential, subscriptionId);
  const result = await client.brokerListener.createOrUpdate(
    "rgiotoperations",
    "resource-name123",
    "resource-name123",
    "resource-name123",
    {
      properties: { ports: [{ port: 1883 }] },
      extendedLocation: {
        name: "qmbrfwcpwwhggszhrdjv",
        type: "CustomLocation",
      },
    },
  );
  console.log(result);
}
