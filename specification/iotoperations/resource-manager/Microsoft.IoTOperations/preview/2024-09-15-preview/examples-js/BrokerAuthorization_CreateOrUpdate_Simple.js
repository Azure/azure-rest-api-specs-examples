const { IoTOperationsClient } = require("@azure/arm-iotoperations");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create a BrokerAuthorizationResource
 *
 * @summary create a BrokerAuthorizationResource
 * x-ms-original-file: 2024-09-15-preview/BrokerAuthorization_CreateOrUpdate_Simple.json
 */
async function brokerAuthorizationCreateOrUpdateSimple() {
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
                clientIds: ["my-client-id"],
                attributes: [{ floor: "floor1", site: "site1" }],
              },
              brokerResources: [
                { method: "Connect" },
                {
                  method: "Subscribe",
                  topics: ["topic", "topic/with/wildcard/#"],
                },
              ],
              stateStoreResources: [{ method: "ReadWrite", keyType: "Pattern", keys: ["*"] }],
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
