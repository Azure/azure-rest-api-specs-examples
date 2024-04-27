const { AzureQuotaExtensionAPI } = require("@azure/arm-quota");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a list of current usage for all resources for the scope specified.
 *
 * @summary Get a list of current usage for all resources for the scope specified.
 * x-ms-original-file: specification/quota/resource-manager/Microsoft.Quota/preview/2023-06-01-preview/examples/getMachineLearningServicesUsages.json
 */
async function quotasListUsagesMachineLearningServices() {
  const scope =
    "subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.MachineLearningServices/locations/eastus";
  const credential = new DefaultAzureCredential();
  const client = new AzureQuotaExtensionAPI(credential);
  const resArray = new Array();
  for await (let item of client.usages.list(scope)) {
    resArray.push(item);
  }
  console.log(resArray);
}
