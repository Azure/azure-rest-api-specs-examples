const { AzureAnalysisServices } = require("@azure/arm-analysisservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Dissociates a Unified Gateway associated with the server.
 *
 * @summary Dissociates a Unified Gateway associated with the server.
 * x-ms-original-file: specification/analysisservices/resource-manager/Microsoft.AnalysisServices/stable/2017-08-01/examples/dissociateGateway.json
 */
async function getDetailsOfAServer() {
  const subscriptionId = "613192d7-503f-477a-9cfe-4efc3ee2bd60";
  const resourceGroupName = "TestRG";
  const serverName = "azsdktest";
  const credential = new DefaultAzureCredential();
  const client = new AzureAnalysisServices(credential, subscriptionId);
  const result = await client.servers.dissociateGateway(resourceGroupName, serverName);
  console.log(result);
}

getDetailsOfAServer().catch(console.error);
