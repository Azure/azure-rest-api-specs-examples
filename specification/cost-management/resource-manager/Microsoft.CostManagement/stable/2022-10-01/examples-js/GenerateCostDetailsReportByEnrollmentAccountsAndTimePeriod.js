const { CostManagementClient } = require("@azure/arm-costmanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to This API is the replacement for all previously release Usage Details APIs. Request to generate a cost details report for the provided date range, billing period (Only enterprise customers) or Invoice Id asynchronously at a certain scope. The initial call to request a report will return a 202 with a 'Location' and 'Retry-After' header. The 'Location' header will provide the endpoint to poll to get the result of the report generation. The 'Retry-After' provides the duration to wait before polling for the generated report. A call to poll the report operation will provide a 202 response with a 'Location' header if the operation is still in progress. Once the report generation operation completes, the polling endpoint will provide a 200 response along with details on the report blob(s) that are available for download. The details on the file(s) available for download will be available in the polling response body. To Understand cost details (formerly known as usage details) fields found in files ,see https://learn.microsoft.com/azure/cost-management-billing/automate/understand-usage-details-fields
 *
 * @summary This API is the replacement for all previously release Usage Details APIs. Request to generate a cost details report for the provided date range, billing period (Only enterprise customers) or Invoice Id asynchronously at a certain scope. The initial call to request a report will return a 202 with a 'Location' and 'Retry-After' header. The 'Location' header will provide the endpoint to poll to get the result of the report generation. The 'Retry-After' provides the duration to wait before polling for the generated report. A call to poll the report operation will provide a 202 response with a 'Location' header if the operation is still in progress. Once the report generation operation completes, the polling endpoint will provide a 200 response along with details on the report blob(s) that are available for download. The details on the file(s) available for download will be available in the polling response body. To Understand cost details (formerly known as usage details) fields found in files ,see https://learn.microsoft.com/azure/cost-management-billing/automate/understand-usage-details-fields
 * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/GenerateCostDetailsReportByEnrollmentAccountsAndTimePeriod.json
 */
async function generateCostDetailsReportByEnrollmentAccountsAndTimePeriod() {
  const scope = "providers/Microsoft.Billing/enrollmentAccounts/1234";
  const parameters = {
    metric: "ActualCost",
    timePeriod: { end: "2020-03-15", start: "2020-03-01" },
  };
  const credential = new DefaultAzureCredential();
  const client = new CostManagementClient(credential);
  const result = await client.generateCostDetailsReport.beginCreateOperationAndWait(
    scope,
    parameters
  );
  console.log(result);
}
