const { AzureQuotaExtensionAPI } = require("@azure/arm-quota");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List all the operations supported by the Microsoft.Quota resource provider.
 *
 * @summary List all the operations supported by the Microsoft.Quota resource provider.
 * x-ms-original-file: specification/quota/resource-manager/Microsoft.Quota/stable/2023-02-01/examples/GetOperations.json
 */
async function getOperations() {
  const credential = new DefaultAzureCredential();
  const client = new AzureQuotaExtensionAPI(credential);
  const resArray = new Array();
  for await (let item of client.quotaOperation.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}
