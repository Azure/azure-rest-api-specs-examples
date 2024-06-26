const { CostManagementClient } = require("@azure/arm-costmanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Generates the reservations details report for provided date range asynchronously by billing profile. The Reservation usage details can be viewed by only certain enterprise roles by default. For more details on the roles see, https://docs.microsoft.com/azure/cost-management-billing/reservations/reservation-utilization#view-utilization-in-the-azure-portal-with-azure-rbac-access
 *
 * @summary Generates the reservations details report for provided date range asynchronously by billing profile. The Reservation usage details can be viewed by only certain enterprise roles by default. For more details on the roles see, https://docs.microsoft.com/azure/cost-management-billing/reservations/reservation-utilization#view-utilization-in-the-azure-portal-with-azure-rbac-access
 * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/GenerateReservationDetailsReportByBillingProfile.json
 */
async function reservationDetails() {
  const billingAccountId = "00000000-0000-0000-0000-000000000000";
  const billingProfileId = "CZSFR-SDFXC-DSDF";
  const startDate = "2020-01-01";
  const endDate = "2020-01-30";
  const credential = new DefaultAzureCredential();
  const client = new CostManagementClient(credential);
  const result = await client.generateReservationDetailsReport.beginByBillingProfileIdAndWait(
    billingAccountId,
    billingProfileId,
    startDate,
    endDate
  );
  console.log(result);
}
