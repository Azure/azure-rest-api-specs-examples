const { CostManagementClient } = require("@azure/arm-costmanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Generates the reservations details report for provided date range asynchronously based on enrollment id. The Reservation usage details can be viewed only by certain enterprise roles. For more details on the roles see, https://docs.microsoft.com/azure/cost-management-billing/manage/understand-ea-roles#usage-and-costs-access-by-role
 *
 * @summary Generates the reservations details report for provided date range asynchronously based on enrollment id. The Reservation usage details can be viewed only by certain enterprise roles. For more details on the roles see, https://docs.microsoft.com/azure/cost-management-billing/manage/understand-ea-roles#usage-and-costs-access-by-role
 * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/GenerateReservationDetailsReportByBillingAccount.json
 */
async function reservationDetails() {
  const billingAccountId = "9845612";
  const startDate = "2020-01-01";
  const endDate = "2020-01-30";
  const credential = new DefaultAzureCredential();
  const client = new CostManagementClient(credential);
  const result = await client.generateReservationDetailsReport.beginByBillingAccountIdAndWait(
    billingAccountId,
    startDate,
    endDate
  );
  console.log(result);
}
