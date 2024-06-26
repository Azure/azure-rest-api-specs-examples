const { CostManagementClient } = require("@azure/arm-costmanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The operation to create or update a view. Update operation requires latest eTag to be set in the request. You may obtain the latest eTag by performing a get operation. Create operation does not require eTag.
 *
 * @summary The operation to create or update a view. Update operation requires latest eTag to be set in the request. You may obtain the latest eTag by performing a get operation. Create operation does not require eTag.
 * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/PrivateViewCreateOrUpdate.json
 */
async function createOrUpdatePrivateView() {
  const viewName = "swaggerExample";
  const parameters = {
    typePropertiesQueryType: "Usage",
    accumulated: "true",
    chart: "Table",
    dataSet: {
      aggregation: { totalCost: { name: "PreTaxCost", function: "Sum" } },
      granularity: "Daily",
      grouping: [],
      sorting: [{ name: "UsageDate", direction: "Ascending" }],
    },
    displayName: "swagger Example",
    eTag: '"1d4ff9fe66f1d10"',
    kpis: [
      { type: "Forecast", enabled: true, id: undefined },
      {
        type: "Budget",
        enabled: true,
        id: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/MYDEVTESTRG/providers/Microsoft.Consumption/budgets/swaggerDemo",
      },
    ],
    metric: "ActualCost",
    pivots: [
      { name: "ServiceName", type: "Dimension" },
      { name: "MeterCategory", type: "Dimension" },
      { name: "swaggerTagKey", type: "TagKey" },
    ],
    timeframe: "MonthToDate",
  };
  const credential = new DefaultAzureCredential();
  const client = new CostManagementClient(credential);
  const result = await client.views.createOrUpdate(viewName, parameters);
  console.log(result);
}
