const { CognitiveServicesManagementClient } = require("@azure/arm-cognitiveservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the private endpoint connections associated with the Cognitive Services account.
 *
 * @summary Gets the private endpoint connections associated with the Cognitive Services account.
 * x-ms-original-file: specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2022-03-01/examples/ListPrivateEndpointConnections.json
 */
async function getPrivateEndpointConnection() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res6977";
  const accountName = "sto2527";
  const credential = new DefaultAzureCredential();
  const client = new CognitiveServicesManagementClient(credential, subscriptionId);
  const result = await client.privateEndpointConnections.list(resourceGroupName, accountName);
  console.log(result);
}

getPrivateEndpointConnection().catch(console.error);
