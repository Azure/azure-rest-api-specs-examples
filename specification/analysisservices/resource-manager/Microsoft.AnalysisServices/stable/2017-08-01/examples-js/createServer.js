const { AzureAnalysisServices } = require("@azure/arm-analysisservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Provisions the specified Analysis Services server based on the configuration specified in the request.
 *
 * @summary Provisions the specified Analysis Services server based on the configuration specified in the request.
 * x-ms-original-file: specification/analysisservices/resource-manager/Microsoft.AnalysisServices/stable/2017-08-01/examples/createServer.json
 */
async function createAServer() {
  const subscriptionId = "613192d7-503f-477a-9cfe-4efc3ee2bd60";
  const resourceGroupName = "TestRG";
  const serverName = "azsdktest";
  const serverParameters = {
    asAdministrators: {
      members: ["azsdktest@microsoft.com", "azsdktest2@microsoft.com"],
    },
    location: "West US",
    sku: { name: "S1", capacity: 1, tier: "Standard" },
    tags: { testKey: "testValue" },
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureAnalysisServices(credential, subscriptionId);
  const result = await client.servers.beginCreateAndWait(
    resourceGroupName,
    serverName,
    serverParameters
  );
  console.log(result);
}

createAServer().catch(console.error);
