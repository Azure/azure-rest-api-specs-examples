const { CognitiveServicesManagementClient } = require("@azure/arm-cognitiveservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Check whether a domain is available.
 *
 * @summary Check whether a domain is available.
 * x-ms-original-file: specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2023-05-01/examples/CheckDomainAvailability.json
 */
async function checkSkuAvailability() {
  const subscriptionId =
    process.env["COGNITIVESERVICES_SUBSCRIPTION_ID"] || "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx";
  const subdomainName = "contosodemoapp1";
  const typeParam = "Microsoft.CognitiveServices/accounts";
  const credential = new DefaultAzureCredential();
  const client = new CognitiveServicesManagementClient(credential, subscriptionId);
  const result = await client.checkDomainAvailability(subdomainName, typeParam);
  console.log(result);
}
