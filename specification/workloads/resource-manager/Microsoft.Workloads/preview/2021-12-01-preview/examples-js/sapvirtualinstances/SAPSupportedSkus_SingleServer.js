const { WorkloadsClient } = require("@azure/arm-workloads");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get SAP supported SKUs.
 *
 * @summary Get SAP supported SKUs.
 * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/preview/2021-12-01-preview/examples/sapvirtualinstances/SAPSupportedSkus_SingleServer.json
 */
async function sapSupportedSkusSingleServer() {
  const subscriptionId = "8e17e36c-42e9-4cd5-a078-7b44883414e0";
  const location = "centralus";
  const credential = new DefaultAzureCredential();
  const client = new WorkloadsClient(credential, subscriptionId);
  const result = await client.sAPSupportedSku(location);
  console.log(result);
}

sapSupportedSkusSingleServer().catch(console.error);
