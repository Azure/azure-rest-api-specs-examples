
/**
 * Samples for GenerateReservationDetailsReport ByBillingProfileId.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/
     * GenerateReservationDetailsReportByBillingProfile.json
     */
    /**
     * Sample code: ReservationDetails.
     * 
     * @param manager Entry point to CostManagementManager.
     */
    public static void reservationDetails(com.azure.resourcemanager.costmanagement.CostManagementManager manager) {
        manager.generateReservationDetailsReports().byBillingProfileId("00000000-0000-0000-0000-000000000000",
            "CZSFR-SDFXC-DSDF", "2020-01-01", "2020-01-30", com.azure.core.util.Context.NONE);
    }
}
