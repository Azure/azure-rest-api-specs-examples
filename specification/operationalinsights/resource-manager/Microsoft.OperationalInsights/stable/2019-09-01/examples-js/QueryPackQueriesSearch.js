const { OperationalInsightsManagementClient } = require("@azure/arm-operationalinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Search a list of Queries defined within a Log Analytics QueryPack according to given search properties.
 *
 * @summary Search a list of Queries defined within a Log Analytics QueryPack according to given search properties.
 * x-ms-original-file: specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2019-09-01/examples/QueryPackQueriesSearch.json
 */
async function querySearch() {
  const subscriptionId = "86dc51d3-92ed-4d7e-947a-775ea79b4918";
  const resourceGroupName = "my-resource-group";
  const queryPackName = "my-querypack";
  const top = 3;
  const includeBody = true;
  const querySearchProperties = {
    related: { categories: ["other", "analytics"] },
    tags: { myLabel: ["label1"] },
  };
  const options = { top, includeBody };
  const credential = new DefaultAzureCredential();
  const client = new OperationalInsightsManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.queries.listSearch(
    resourceGroupName,
    queryPackName,
    querySearchProperties,
    options
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

querySearch().catch(console.error);
