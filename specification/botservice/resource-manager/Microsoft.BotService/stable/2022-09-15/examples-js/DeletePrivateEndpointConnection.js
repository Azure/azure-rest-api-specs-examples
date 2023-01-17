const { AzureBotService } = require("@azure/arm-botservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the specified private endpoint connection associated with the Bot.
 *
 * @summary Deletes the specified private endpoint connection associated with the Bot.
 * x-ms-original-file: specification/botservice/resource-manager/Microsoft.BotService/stable/2022-09-15/examples/DeletePrivateEndpointConnection.json
 */
async function deletePrivateEndpointConnection() {
  const subscriptionId = process.env["BOTSERVICE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const resourceGroupName = process.env["BOTSERVICE_RESOURCE_GROUP"] || "res6977";
  const resourceName = "sto2527";
  const privateEndpointConnectionName = "{privateEndpointConnectionName}";
  const credential = new DefaultAzureCredential();
  const client = new AzureBotService(credential, subscriptionId);
  const result = await client.privateEndpointConnections.delete(
    resourceGroupName,
    resourceName,
    privateEndpointConnectionName
  );
  console.log(result);
}
