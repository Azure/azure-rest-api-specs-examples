const { AzureAnalysisServices } = require("@azure/arm-analysisservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists eligible SKUs for an Analysis Services resource.
 *
 * @summary Lists eligible SKUs for an Analysis Services resource.
 * x-ms-original-file: specification/analysisservices/resource-manager/Microsoft.AnalysisServices/stable/2017-08-01/examples/listSkusForExisting.json
 */
async function listEligibleSkUsForAnExistingServer() {
  const subscriptionId = "613192d7-503f-477a-9cfe-4efc3ee2bd60";
  const resourceGroupName = "TestRG";
  const serverName = "azsdktest";
  const credential = new DefaultAzureCredential();
  const client = new AzureAnalysisServices(credential, subscriptionId);
  const result = await client.servers.listSkusForExisting(resourceGroupName, serverName);
  console.log(result);
}

listEligibleSkUsForAnExistingServer().catch(console.error);
