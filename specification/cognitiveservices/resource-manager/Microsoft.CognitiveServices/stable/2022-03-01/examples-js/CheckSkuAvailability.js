const { CognitiveServicesManagementClient } = require("@azure/arm-cognitiveservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Check available SKUs.
 *
 * @summary Check available SKUs.
 * x-ms-original-file: specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2022-03-01/examples/CheckSkuAvailability.json
 */
async function checkSkuAvailability() {
  const subscriptionId = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx";
  const location = "westus";
  const skus = ["S0"];
  const kind = "Face";
  const typeParam = "Microsoft.CognitiveServices/accounts";
  const credential = new DefaultAzureCredential();
  const client = new CognitiveServicesManagementClient(credential, subscriptionId);
  const result = await client.checkSkuAvailability(location, skus, kind, typeParam);
  console.log(result);
}

checkSkuAvailability().catch(console.error);
