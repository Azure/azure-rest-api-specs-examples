const { WorkloadsClient } = require("@azure/arm-workloads");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get SAP sizing recommendations by providing input SAPS for application tier and memory required for database tier
 *
 * @summary Get SAP sizing recommendations by providing input SAPS for application tier and memory required for database tier
 * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/stable/2023-04-01/examples/sapvirtualinstances/SAPSizingRecommendations_S4HANA_HA_AvZone.json
 */
async function sapSizingRecommendationsS4HanaDistributedHaAvZone() {
  const subscriptionId =
    process.env["WORKLOADS_SUBSCRIPTION_ID"] || "8e17e36c-42e9-4cd5-a078-7b44883414e0";
  const location = "centralus";
  const credential = new DefaultAzureCredential();
  const client = new WorkloadsClient(credential, subscriptionId);
  const result = await client.sAPSizingRecommendations(location);
  console.log(result);
}
