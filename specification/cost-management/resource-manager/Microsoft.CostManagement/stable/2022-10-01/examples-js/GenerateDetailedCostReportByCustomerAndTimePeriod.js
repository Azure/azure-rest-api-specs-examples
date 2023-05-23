const { CostManagementClient } = require("@azure/arm-costmanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Generates the detailed cost report for provided date range, billing period(only enterprise customers) or Invoice ID asynchronously at a certain scope. Call returns a 202 with header Azure-Consumption-AsyncOperation providing a link to the operation created. A call on the operation will provide the status and if the operation is completed the blob file where generated detailed cost report is being stored.
 *
 * @summary Generates the detailed cost report for provided date range, billing period(only enterprise customers) or Invoice ID asynchronously at a certain scope. Call returns a 202 with header Azure-Consumption-AsyncOperation providing a link to the operation created. A call on the operation will provide the status and if the operation is completed the blob file where generated detailed cost report is being stored.
 * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/GenerateDetailedCostReportByCustomerAndTimePeriod.json
 */
async function generateDetailedCostReportByCustomerAndTimePeriod() {
  const scope = "providers/Microsoft.Billing/billingAccounts/12345:6789/customers/13579";
  const parameters = {
    metric: "ActualCost",
    timePeriod: { end: "2020-03-15", start: "2020-03-01" },
  };
  const credential = new DefaultAzureCredential();
  const client = new CostManagementClient(credential);
  const result = await client.generateDetailedCostReport.beginCreateOperationAndWait(
    scope,
    parameters
  );
  console.log(result);
}
