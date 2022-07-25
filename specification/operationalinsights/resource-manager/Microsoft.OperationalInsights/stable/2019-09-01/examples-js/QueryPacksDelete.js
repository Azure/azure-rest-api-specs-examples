const { OperationalInsightsManagementClient } = require("@azure/arm-operationalinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a Log Analytics QueryPack.
 *
 * @summary Deletes a Log Analytics QueryPack.
 * x-ms-original-file: specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2019-09-01/examples/QueryPacksDelete.json
 */
async function queryPacksDelete() {
  const subscriptionId = "86dc51d3-92ed-4d7e-947a-775ea79b4919";
  const resourceGroupName = "my-resource-group";
  const queryPackName = "my-querypack";
  const credential = new DefaultAzureCredential();
  const client = new OperationalInsightsManagementClient(credential, subscriptionId);
  const result = await client.queryPacks.delete(resourceGroupName, queryPackName);
  console.log(result);
}

queryPacksDelete().catch(console.error);
