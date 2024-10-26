const { IoTOperationsClient } = require("@azure/arm-iotoperations");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create a BrokerAuthorizationResource
 *
 * @summary create a BrokerAuthorizationResource
 * x-ms-original-file: 2024-09-15-preview/BrokerAuthorization_CreateOrUpdate_Complex.json
 */
async function brokerAuthorizationCreateOrUpdateComplex() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "F8C729F9-DF9C-4743-848F-96EE433D8E53";
  const client = new IoTOperationsClient(credential, subscriptionId);
  const result = await client.brokerAuthorization.createOrUpdate(
    "rgiotoperations",
    "resource-name123",
    "resource-name123",
    "resource-name123",
    {
      properties: {
        authorizationPolicies: {
          cache: "Enabled",
          rules: [
            {
              principals: {
                usernames: ["temperature-sensor", "humidity-sensor"],
                attributes: [{ building: "17", organization: "contoso" }],
              },
              brokerResources: [
                {
                  method: "Connect",
                  clientIds: ["{principal.attributes.building}*"],
                },
                {
                  method: "Publish",
                  topics: [
                    "sensors/{principal.attributes.building}/{principal.clientId}/telemetry/*",
                  ],
                },
                {
                  method: "Subscribe",
                  topics: ["commands/{principal.attributes.organization}"],
                },
              ],
              stateStoreResources: [
                {
                  method: "Read",
                  keyType: "Pattern",
                  keys: [
                    "myreadkey",
                    "myotherkey?",
                    "mynumerickeysuffix[0-9]",
                    "clients:{principal.clientId}:*",
                  ],
                },
                {
                  method: "ReadWrite",
                  keyType: "Binary",
                  keys: ["MTE2IDEwMSAxMTUgMTE2"],
                },
              ],
            },
          ],
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
