const { IoTOperationsClient } = require("@azure/arm-iotoperations");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create a BrokerAuthorizationResource
 *
 * @summary create a BrokerAuthorizationResource
 * x-ms-original-file: 2024-09-15-preview/BrokerAuthorization_CreateOrUpdate_MaximumSet_Gen.json
 */
async function brokerAuthorizationCreateOrUpdate() {
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
              brokerResources: [{ method: "Connect", clientIds: ["nlc"], topics: ["wvuca"] }],
              principals: {
                attributes: [{ key5526: "nydhzdhbldygqcn" }],
                clientIds: ["smopeaeddsygz"],
                usernames: ["iozngyqndrteikszkbasinzdjtm"],
              },
              stateStoreResources: [
                {
                  keyType: "Pattern",
                  keys: ["tkounsqtwvzyaklxjqoerpu"],
                  method: "Read",
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
