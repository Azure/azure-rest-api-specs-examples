const { WorkloadsClient } = require("@azure/arm-workloads");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get SAP sizing recommendations.
 *
 * @summary Get SAP sizing recommendations.
 * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/preview/2021-12-01-preview/examples/sapvirtualinstances/SAPSizingRecommendations_S4HANA_SingleServer.json
 */
async function sapSizingRecommendationsS4HanaSingleServer() {
  const subscriptionId = "8e17e36c-42e9-4cd5-a078-7b44883414e0";
  const location = "centralus";
  const credential = new DefaultAzureCredential();
  const client = new WorkloadsClient(credential, subscriptionId);
  const result = await client.sAPSizingRecommendations(location);
  console.log(result);
}

sapSizingRecommendationsS4HanaSingleServer().catch(console.error);
