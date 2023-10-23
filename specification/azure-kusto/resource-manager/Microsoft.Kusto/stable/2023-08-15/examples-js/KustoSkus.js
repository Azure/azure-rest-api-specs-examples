const { KustoManagementClient } = require("@azure/arm-kusto");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists eligible region SKUs for Kusto resource provider by Azure region.
 *
 * @summary Lists eligible region SKUs for Kusto resource provider by Azure region.
 * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2023-08-15/examples/KustoSkus.json
 */
async function kustoListRegionSkus() {
  const subscriptionId =
    process.env["KUSTO_SUBSCRIPTION_ID"] || "12345678-1234-1234-1234-123456789098";
  const location = "westus";
  const credential = new DefaultAzureCredential();
  const client = new KustoManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.skus.list(location)) {
    resArray.push(item);
  }
  console.log(resArray);
}
