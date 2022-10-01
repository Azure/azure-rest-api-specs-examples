const { AzureAnalysisServices } = require("@azure/arm-analysisservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Check the name availability in the target location.
 *
 * @summary Check the name availability in the target location.
 * x-ms-original-file: specification/analysisservices/resource-manager/Microsoft.AnalysisServices/stable/2017-08-01/examples/checkNameAvailability.json
 */
async function getDetailsOfAServer() {
  const subscriptionId = "613192d7-503f-477a-9cfe-4efc3ee2bd60";
  const location = "West US";
  const serverParameters = {
    name: "azsdktest",
    type: "Microsoft.AnalysisServices/servers",
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureAnalysisServices(credential, subscriptionId);
  const result = await client.servers.checkNameAvailability(location, serverParameters);
  console.log(result);
}

getDetailsOfAServer().catch(console.error);
