const { AzureBotService } = require("@azure/arm-botservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Register a new Auth Connection for a Bot Service
 *
 * @summary Register a new Auth Connection for a Bot Service
 * x-ms-original-file: specification/botservice/resource-manager/Microsoft.BotService/preview/2021-05-01-preview/examples/PutConnection.json
 */
async function createConnectionSetting() {
  const subscriptionId = "subscription-id";
  const resourceGroupName = "OneResourceGroupName";
  const resourceName = "samplebotname";
  const connectionName = "sampleConnection";
  const parameters = {
    etag: "etag1",
    location: "West US",
    properties: {
      clientId: "sampleclientid",
      clientSecret: "samplesecret",
      parameters: [
        { key: "key1", value: "value1" },
        { key: "key2", value: "value2" },
      ],
      scopes: "samplescope",
      serviceProviderId: "serviceproviderid",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureBotService(credential, subscriptionId);
  const result = await client.botConnection.create(
    resourceGroupName,
    resourceName,
    connectionName,
    parameters
  );
  console.log(result);
}

createConnectionSetting().catch(console.error);
