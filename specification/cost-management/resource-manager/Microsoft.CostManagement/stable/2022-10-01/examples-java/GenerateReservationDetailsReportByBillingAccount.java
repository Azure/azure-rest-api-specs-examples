/** Samples for GenerateReservationDetailsReport ByBillingAccountId. */
public final class Main {
    /*
     * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/GenerateReservationDetailsReportByBillingAccount.json
     */
    /**
     * Sample code: ReservationDetails.
     *
     * @param manager Entry point to CostManagementManager.
     */
    public static void reservationDetails(com.azure.resourcemanager.costmanagement.CostManagementManager manager) {
        manager
            .generateReservationDetailsReports()
            .byBillingAccountId("9845612", "2020-01-01", "2020-01-30", com.azure.core.util.Context.NONE);
    }
}
