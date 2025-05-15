const { CarbonOptimizationManagementClient } = require("@azure/arm-carbonoptimization");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to aPI for Carbon Emissions Reports
 *
 * @summary aPI for Carbon Emissions Reports
 * x-ms-original-file: 2025-04-01/queryCarbonEmissionsTopNResourceGroupItemsReport.json
 */
async function queryCarbonEmissionTopNResourceGroupReport() {
  const credential = new DefaultAzureCredential();
  const client = new CarbonOptimizationManagementClient(credential);
  const result = await client.carbonService.queryCarbonEmissionReports({
    reportType: "TopItemsSummaryReport",
    subscriptionList: [
      "00000000-0000-0000-0000-000000000000",
      "00000000-0000-0000-0000-000000000001,",
      "00000000-0000-0000-0000-000000000002",
      "00000000-0000-0000-0000-000000000003",
      "00000000-0000-0000-0000-000000000004",
      "00000000-0000-0000-0000-000000000005",
      "00000000-0000-0000-0000-000000000006",
      "00000000-0000-0000-0000-000000000007",
      "00000000-0000-0000-0000-000000000008",
    ],
    carbonScopeList: ["Scope1", "Scope3"],
    dateRange: { start: "2024-05-01", end: "2024-05-01" },
    categoryType: "ResourceGroup",
    topItems: 5,
  });
  console.log(result);
}
