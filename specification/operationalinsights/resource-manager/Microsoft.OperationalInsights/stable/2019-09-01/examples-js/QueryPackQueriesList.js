const { OperationalInsightsManagementClient } = require("@azure/arm-operationalinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a list of Queries defined within a Log Analytics QueryPack.
 *
 * @summary Gets a list of Queries defined within a Log Analytics QueryPack.
 * x-ms-original-file: specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2019-09-01/examples/QueryPackQueriesList.json
 */
async function queryList() {
  const subscriptionId = "86dc51d3-92ed-4d7e-947a-775ea79b4918";
  const resourceGroupName = "my-resource-group";
  const queryPackName = "my-querypack";
  const includeBody = true;
  const options = { includeBody };
  const credential = new DefaultAzureCredential();
  const client = new OperationalInsightsManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.queries.list(resourceGroupName, queryPackName, options)) {
    resArray.push(item);
  }
  console.log(resArray);
}

queryList().catch(console.error);
