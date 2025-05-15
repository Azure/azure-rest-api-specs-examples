const { CarbonOptimizationManagementClient } = require("@azure/arm-carbonoptimization");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to aPI for Carbon Emissions Reports
 *
 * @summary aPI for Carbon Emissions Reports
 * x-ms-original-file: 2025-04-01/queryCarbonEmissionsMonthlySummaryReportWithOtherOptionalFilter.json
 */
async function queryCarbonEmissionMonthlySummaryReportWithOptionalFilterLocationListResourceTypeListResourceGroupUrlList() {
  const credential = new DefaultAzureCredential();
  const client = new CarbonOptimizationManagementClient(credential);
  const result = await client.carbonService.queryCarbonEmissionReports({
    reportType: "MonthlySummaryReport",
    subscriptionList: ["00000000-0000-0000-0000-000000000000"],
    carbonScopeList: ["Scope1", "Scope3"],
    dateRange: { start: "2024-03-01", end: "2024-05-01" },
    locationList: ["east us", "west us"],
    resourceTypeList: ["microsoft.storage/storageaccounts", "microsoft.databricks/workspaces"],
    resourceGroupUrlList: [
      "/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg-name",
    ],
  });
  console.log(result);
}
