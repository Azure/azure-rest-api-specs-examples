const { AzureBotService } = require("@azure/arm-botservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update the state of specified private endpoint connection associated with the Bot.
 *
 * @summary Update the state of specified private endpoint connection associated with the Bot.
 * x-ms-original-file: specification/botservice/resource-manager/Microsoft.BotService/stable/2022-09-15/examples/PutPrivateEndpointConnection.json
 */
async function putPrivateEndpointConnection() {
  const subscriptionId = process.env["BOTSERVICE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const resourceGroupName = process.env["BOTSERVICE_RESOURCE_GROUP"] || "res7687";
  const resourceName = "sto9699";
  const privateEndpointConnectionName = "{privateEndpointConnectionName}";
  const properties = {
    privateLinkServiceConnectionState: {
      description: "Auto-Approved",
      status: "Approved",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureBotService(credential, subscriptionId);
  const result = await client.privateEndpointConnections.create(
    resourceGroupName,
    resourceName,
    privateEndpointConnectionName,
    properties
  );
  console.log(result);
}
